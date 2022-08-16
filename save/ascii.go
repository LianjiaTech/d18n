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

	"github.com/olekukonko/tablewriter"
)

// saveRows2ASCII print rows as ascii table
func saveRows2ASCII(s *SaveStruct, rows *sql.Rows) error {
	table := tablewriter.NewWriter(os.Stdout)

	// column info
	columnNames, err := rows.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	// set table header
	if !s.Config.NoHeader && !s.Config.Vertical {
		table.SetHeader(columnNames)
	}

	// init columns
	columnValues := make([]interface{}, len(columnNames))
	cols := make([]interface{}, len(columnNames))
	for j := range columnValues {
		cols[j] = &columnValues[j]
	}

	// set every rows
	for rows.Next() {
		s.Status.Lines++
		// preview only show first N lines
		if s.Config.Preview != 0 && s.Status.Lines > s.Config.Preview {
			break
		}

		// limit return rows
		if s.Config.Limit != 0 && s.Status.Lines > s.Config.Limit {
			s.Status.Lines = s.Config.Limit
			break
		}

		// scan columns
		if err := rows.Scan(cols...); err != nil {
			return err
		}

		values := make([]string, len(columnNames))
		for j, col := range columnValues {
			if col == nil {
				values[j] = s.Config.NULLString
			} else {
				values[j] = s.String(col, columnTypes[j])

				// data mask
				values[j], err = s.Masker.Mask(s.FieldName(j), values[j])
				if err != nil {
					return err
				}

				// hex-blob
				values[j], _ = s.Config.Hex(s.FieldName(j), values[j])
			}
		}
		if s.Config.Vertical {
			table.Append([]string{fmt.Sprintf(`********* Row %d *********`, s.Status.Lines), fmt.Sprintf(`********* Row %d *********`, s.Status.Lines)})
			for i, v := range values {
				table.Append([]string{columnNames[i], v})
			}
		} else {
			table.Append(values)
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	// print table
	if len(columnNames) > 0 { // `do 1` only return empty set without columns
		table.Render()
	}
	return nil
}
