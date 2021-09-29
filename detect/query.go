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
	"d18n/common"
	"database/sql"
	"regexp"
)

type DetectStruct struct {
	CommonConfig common.Config
	Status       DetectStatus
}

func NewDetectStruct(c common.Config) (*DetectStruct, error) {
	var d *DetectStruct
	err := ParseSensitiveConfig(common.Cfg.Sensitive)
	if err != nil {
		return d, err
	}

	d = &DetectStruct{
		CommonConfig: c,
		Status: DetectStatus{
			Columns: make(map[string][]string),
		},
	}
	return d, nil
}

// DetectQuery check data from query result
func (d *DetectStruct) DetectQuery() error {

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
	d.Status.Header = common.DBParseColumnTypes(header)

	// check column names
	checkQueryHeader(d)

	// check column values
	for rows.Next() {
		d.Status.Lines++
		// limit return rows
		if d.CommonConfig.Limit != 0 && d.Status.Lines > d.CommonConfig.Limit {
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
				d.Status.Columns[d.Status.Header[j].Name] = append(d.Status.Columns[d.Status.Header[j].Name], checkValue(col.String)...)
			}
		}
	}

	err = rows.Err()

	return err
}

func checkQueryHeader(d *DetectStruct) {

	d.Status.Columns = make(map[string][]string, len(d.Status.Header))

	for _, h := range d.Status.Header {
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
		d.Status.Columns[key] = types
	}
}
