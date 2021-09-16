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
	"d18n/mask"

	"github.com/olekukonko/tablewriter"
)

// saveRows2ASCII print rows as ascii table
func saveRows2ASCII(rows *sql.Rows) error {
	var err error

	table := tablewriter.NewWriter(os.Stdout)

	// set table header
	if !common.Cfg.NoHeader {
		table.SetHeader(common.DBParserColumnNames(saveStatus.Header))
	}

	// init columns
	columns := make([]interface{}, len(saveStatus.Header))
	cols := make([]interface{}, len(saveStatus.Header))
	for j := range columns {
		cols[j] = &columns[j]
	}

	// set every rows
	for rows.Next() {
		saveStatus.Lines++
		// preview only show first N lines
		if common.Cfg.Preview != 0 && saveStatus.Lines > common.Cfg.Preview {
			break
		}

		// limit return rows
		if common.Cfg.Limit != 0 && saveStatus.Lines > common.Cfg.Limit {
			break
		}

		// scan columns
		if err := rows.Scan(cols...); err != nil {
			return err
		}

		values := make([]string, len(columns))
		for j, col := range columns {
			if col == nil {
				values[j] = common.Cfg.NULLString
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
				values[j], err = mask.Mask(saveStatus.Header[j].Name(), values[j])
				if err != nil {
					return err
				}

				// hex-blob
				values[j], _ = common.HexBLOB(saveStatus.Header[j].Name(), values[j])
			}
		}
		table.Append(values)
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	// print table
	if len(saveStatus.Header) > 0 { // `do 1` only return empty set without columns
		table.Render()
	}
	return nil
}
