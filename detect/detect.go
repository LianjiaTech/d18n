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
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"d18n/common"
)

// BasicDetect basic sensitive data detection regexp check rules
type BasicDetect struct {
	Key   []string `yaml:"key"`   // check column name by regexp
	Value []string `yaml:"value"` // check column value by regexp
}

// DetectStatus ...
type DetectStatus struct {
	Header        []common.HeaderColumn // Column Header, include column type info
	Columns       map[string][]string   // detected sensitive columns
	Lines         int                   // lines checked
	QueryTimeCost int64                 // query time cost
	TimeCost      int64                 // time cost
}

type DetectStruct struct {
	Config    common.Config
	Sensitive sensitiveConfig
	Status    DetectStatus
}

func NewDetectStruct(c common.Config) (*DetectStruct, error) {
	var d = &DetectStruct{
		Config: c,
		Status: DetectStatus{
			Columns: make(map[string][]string),
		},
	}

	err := d.parseConfig()
	if err != nil {
		return d, err
	}

	return d, nil
}

// Detect sensitive data detection
func (d *DetectStruct) Detect() error {
	var err error
	detectStartTime := time.Now().UnixNano()

	d.Status.Columns = make(map[string][]string)

	switch d.Config.File {
	case "stdout", "":
		if d.Config.Query != "" {
			err = d.DetectQuery()
		} else {
			return fmt.Errorf("no data source or file to check")
		}
	default:
		err = d.DetectFile()
	}

	detectEndTime := time.Now().UnixNano()
	d.Status.TimeCost = detectEndTime - detectStartTime

	return err
}

// DetectQuery check data from query result
func (d *DetectStruct) DetectQuery() error {

	rows, err := d.Config.QueryRows()
	if err != nil {
		return err
	}
	defer rows.Close()

	// get column header
	header, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	d.Status.Header = d.Config.DBParseColumnTypes(header)

	// check column names
	d.checkHeader()

	// check column values
	for rows.Next() {
		d.Status.Lines++
		// limit return rows
		if d.Config.Limit != 0 && d.Status.Lines > d.Config.Limit {
			break
		}

		columns := make([]sql.NullString, len(header))
		cols := make([]interface{}, len(header))
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
				d.Status.Columns[d.Status.Header[j].Name] = append(d.Status.Columns[d.Status.Header[j].Name], d.checkValue(col.String)...)
			}
		}
	}

	err = rows.Err()

	return err
}

// DetectFile check data from file
func (d *DetectStruct) DetectFile() error {
	var err error

	// get column header
	if d.Config.Schema != "" {
		d.Status.Header, err = d.Config.TableTemplate()
		if err != nil {
			return err
		}
	}

	// file type switch
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(d.Config.File), "."))
	switch suffix {
	case "tsv": // tab-separated values
		d.Config.Comma = '\t'
		err = d.detectCSV()
	case "txt": // space-separated values
		d.Config.Comma = ' '
		err = d.detectCSV()
	case "psv": // pipe-separated values
		d.Config.Comma = '|'
		err = d.detectCSV()
	case "csv": // comma-separated values
		d.Config.Comma = ','
		err = d.detectCSV()
	case "xlsx": // microsoft office excel
		err = d.detectXlsx()
	case "html": // html format
		err = d.detectHTML()
	case "sql": // sql file
		err = d.detectSQL()
	case "json": // json file
		err = d.detectJSON()
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}
	return err
}

func (d *DetectStruct) checkHeader() {
	d.Status.Columns = make(map[string][]string, len(d.Status.Header))

	for _, h := range d.Status.Header {
		key := h.Name
		var types []string

		// sensitive key word check
		for t, rule := range d.Sensitive {
			for _, k := range rule.Key {
				r := regexp.MustCompile(k)
				if r.MatchString(key) {
					types = append(types, t)
					break
				}
			}
		}

		// update d.Status.Columns
		d.Status.Columns[key] = types
	}
}

func (d *DetectStruct) checkValue(value string) []string {
	var types []string

	// only check first 256 bytes, too long sentence may contain any thing
	if len(value) > 256 {
		value = value[:256]
	}

	// regexp
	for t, rule := range d.Sensitive {
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

// ShowStatus print detect status
func (d *DetectStruct) ShowStatus() error {
	for k, v := range d.Status.Columns {
		d.Status.Columns[k] = common.StringUnique(v)
	}

	s, err := json.MarshalIndent(d.Status.Columns, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(s))

	// verbose mode
	if !d.Config.Verbose {
		return nil
	}

	println(
		"Detect Rows:", d.Status.Lines, "Total Cost:", fmt.Sprint(time.Duration(d.Status.TimeCost)*time.Nanosecond),
	)
	return nil
}
