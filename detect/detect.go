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

// sensitiveConfig ...
var sensitiveConfig Config

// DetectStatus ...
type DetectStatus struct {
	Header        []common.HeaderColumn // Column Header, include column type info
	Columns       map[string][]string   // detected sensitive columns
	Lines         int                   // lines checked
	QueryTimeCost int64                 // query time cost
	TimeCost      int64                 // time cost
}

func ParseSensitiveConfig(file string) error {
	// load sensitive config
	buf, err := ioutil.ReadFile(file)
	if err == nil {
		defaultSensitiveConfig = buf
	}

	err = yaml.Unmarshal(defaultSensitiveConfig, &sensitiveConfig)
	if err != nil {
		return err
	}

	// check config regexp valid
	for _, v := range sensitiveConfig {
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
func (d *DetectStruct) Detect() error {
	var err error
	detectStartTime := time.Now().UnixNano()

	d.Status.Columns = make(map[string][]string)

	switch d.CommonConfig.File {
	case "stdout", "":
		if d.CommonConfig.Query != "" {
			//d, err := NewDetectStruct(d.CommonConfig)
			//if err != nil {
			//	return err
			//}
			err = d.DetectQuery()
			//d.Status = d.Status
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

// DetectFile check data from file
func (d *DetectStruct) DetectFile() error {
	var err error

	// get column header
	if d.CommonConfig.Schema != "" {
		d.Status.Header, err = common.TableTemplate()
		if err != nil {
			return err
		}
	}

	// file type switch
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(d.CommonConfig.File), "."))
	switch suffix {
	case "tsv": // tab-separated values
		d.CommonConfig.Comma = '\t'
		err = d.detectCSV()
	case "txt": // space-separated values
		d.CommonConfig.Comma = ' '
		err = d.detectCSV()
	case "psv": // pipe-separated values
		d.CommonConfig.Comma = '|'
		err = d.detectCSV()
	case "csv": // comma-separated values
		d.CommonConfig.Comma = ','
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

func checkFileHeader(detectStatus DetectStatus, header []common.HeaderColumn) {
	detectStatus.Columns = make(map[string][]string, len(header))

	for _, h := range header {
		key := h.Name
		var types []string

		// sensitive key word check
		for t, rule := range sensitiveConfig {
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
	for t, rule := range sensitiveConfig {
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
func (d *DetectStruct) CheckStatus() error {
	for k, v := range d.Status.Columns {
		d.Status.Columns[k] = common.StringUnique(v)
	}

	s, err := json.MarshalIndent(d.Status.Columns, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(s))

	// verbose mode
	if !d.CommonConfig.Verbose {
		return nil
	}

	println(
		"Detect Rows:", d.Status.Lines, "Total Cost:", fmt.Sprint(time.Duration(d.Status.TimeCost)*time.Nanosecond),
	)
	return nil
}
