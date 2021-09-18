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

package save

import (
	"database/sql"
	"fmt"
	"os"

	"d18n/common"

	"github.com/olekukonko/tablewriter"
)

// saveRows2ASCII print rows as ascii table
func saveRows2ASCII(s *SaveStruct, rows *sql.Rows) error {
	var err error

	table := tablewriter.NewWriter(os.Stdout)

	// set table header
	if !s.CommonConfig.NoHeader {
		table.SetHeader(common.DBParserColumnNames(s.Status.Header))
	}

	// init columns
	columns := make([]interface{}, len(s.Status.Header))
	cols := make([]interface{}, len(s.Status.Header))
	for j := range columns {
		cols[j] = &columns[j]
	}

	// set every rows
	for rows.Next() {
		s.Status.Lines++
		// preview only show first N lines
		if s.CommonConfig.Preview != 0 && s.Status.Lines > s.CommonConfig.Preview {
			break
		}

		// limit return rows
		if s.CommonConfig.Limit != 0 && s.Status.Lines > s.CommonConfig.Limit {
			break
		}

		// scan columns
		if err := rows.Scan(cols...); err != nil {
			return err
		}

		values := make([]string, len(columns))
		for j, col := range columns {
			if col == nil {
				values[j] = s.CommonConfig.NULLString
			} else {
				switch col.(type) {
				case []byte:
					values[j] = string(col.([]byte))
				case []string:
					values[j] = common.ParseArray(col.([]string))
				default:
					values[j] = fmt.Sprint(col)
				}

				// data mask
				values[j], err = s.Masker.Mask(s.Status.Header[j].Name(), values[j])
				if err != nil {
					return err
				}

				// hex-blob
				values[j], _ = common.HexBLOB(s.Status.Header[j].Name(), values[j])
			}
		}
		table.Append(values)
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	// print table
	if len(s.Status.Header) > 0 { // `do 1` only return empty set without columns
		table.Render()
	}
	return nil
}
