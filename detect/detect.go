/*
 * Copyright(c)  2021 Lianjia, Inc.  All Rights Reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package detect

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"d18n/common"

	"gopkg.in/yaml.v2"
)

//go:embed sensitive.yaml
var defaultSensitiveConfig []byte

// BasicDetect basic sensitive data detection regexp check rules
type BasicDetect struct {
	Key   []string `yaml:"key"`   // check column name by regexp
	Value []string `yaml:"value"` // check column value by regexp
}

// Config sensitive config
type Config map[string]BasicDetect

// SensitiveConfig ...
var SensitiveConfig Config

// DetectStatus ...
type DetectStatus struct {
	Header        []common.HeaderColumn // Column Header, include column type info
	Columns       map[string][]string   // detected sensitive columns
	Lines         int                   // lines checked
	QueryTimeCost int64                 // query time cost
	TimeCost      int64                 // time cost
}

var detectStatus DetectStatus

func ParseSensitiveConfig() error {
	// load sensitive config
	buf, err := ioutil.ReadFile(common.Cfg.Sensitive)
	if err == nil {
		defaultSensitiveConfig = buf
	}

	err = yaml.Unmarshal(defaultSensitiveConfig, &SensitiveConfig)
	if err != nil {
		return err
	}

	// check config regexp valid
	for _, v := range SensitiveConfig {
		for _, r := range v.Key {
			_, err = regexp.Compile(r)
			if err != nil {
				return err
			}
		}

		for _, r := range v.Value {
			_, err = regexp.Compile(r)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Detect sensitive data detection
func Detect() error {
	var err error
	detectStartTime := time.Now().UnixNano()

	detectStatus.Columns = make(map[string][]string)

	switch common.Cfg.File {
	case "stdout", "":
		if common.Cfg.Query != "" {
			err = detectFromQuery()
		} else {
			return fmt.Errorf("no data source or file to check")
		}
	default:
		err = detectFromFile()
	}

	detectEndTime := time.Now().UnixNano()
	detectStatus.TimeCost = detectEndTime - detectStartTime

	return err
}

// detectFromFile check data from file
func detectFromFile() error {
	var err error

	// get column header
	if common.Cfg.Schema != "" {
		detectStatus.Header, err = common.TableTemplate()
		if err != nil {
			return err
		}
	}

	// file type switch
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(common.Cfg.File), "."))
	switch suffix {
	case "tsv": // tab-separated values
		common.Cfg.Comma = '\t'
		err = detectCSV()
	case "txt": // space-separated values
		common.Cfg.Comma = ' '
		err = detectCSV()
	case "psv": // pipe-separated values
		common.Cfg.Comma = '|'
		err = detectCSV()
	case "csv": // comma-separated values
		common.Cfg.Comma = ','
		err = detectCSV()
	case "xlsx": // microsoft office excel
		err = detectXlsx()
	case "html": // html format
		err = detectHTML()
	case "sql": // sql file
		err = detectSQL()
	case "json": // json file
		err = detectJSON()
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}
	return err
}

// detectFromQuery check data from query result
func detectFromQuery() error {

	rows, err := common.QueryRows()
	if err != nil {
		return err
	}
	defer rows.Close()

	// get column header
	header, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	detectStatus.Header = common.DBParseColumnTypes(header)

	// check column names
	checkHeader(detectStatus.Header)

	// check column values
	for rows.Next() {
		detectStatus.Lines++
		// limit return rows
		if common.Cfg.Limit != 0 && detectStatus.Lines > common.Cfg.Limit {
			break
		}

		columns := make([]sql.NullString, len(detectStatus.Header))
		cols := make([]interface{}, len(detectStatus.Header))
		for j := range columns {
			cols[j] = &columns[j]
		}

		if err := rows.Scan(cols...); err != nil {
			return err
		}

		// check value
		for j, col := range columns {
			// pass NULL string
			if col.Valid {
				detectStatus.Columns[detectStatus.Header[j].Name] = append(detectStatus.Columns[detectStatus.Header[j].Name], checkValue(col.String)...)
			}
		}
	}

	err = rows.Err()

	return err
}

func checkHeader(header []common.HeaderColumn) {

	detectStatus.Columns = make(map[string][]string, len(header))

	for _, h := range header {
		key := h.Name
		var types []string

		// sensitive key word check
		for t, rule := range SensitiveConfig {
			for _, k := range rule.Key {
				r := regexp.MustCompile(k)
				if r.MatchString(key) {
					types = append(types, t)
					break
				}
			}
		}

		// update detectStatus.Columns
		detectStatus.Columns[key] = types
	}
}

func checkValue(value string) []string {
	var types []string

	// only check first 256 bytes, too long sentence may contain any thing
	if len(value) > 256 {
		value = value[:256]
	}

	// regexp
	for t, rule := range SensitiveConfig {
		for _, k := range rule.Value {
			r := regexp.MustCompile(k)
			if r.MatchString(value) {
				types = append(types, t)
				break
			}
		}
	}

	// NLP gse
	if t := GSE(value); t != "" {
		types = append(types, t)
	}

	return types
}

// CheckStatus print detect status
func CheckStatus() error {
	for k, v := range detectStatus.Columns {
		detectStatus.Columns[k] = common.StringUnique(v)
	}

	s, err := json.MarshalIndent(detectStatus.Columns, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(s))

	// verbose mode
	if !common.Cfg.Verbose {
		return nil
	}

	println(
		"Detect Rows:", detectStatus.Lines, "Total Cost:", fmt.Sprint(time.Duration(detectStatus.TimeCost)*time.Nanosecond),
	)
	return nil
}
