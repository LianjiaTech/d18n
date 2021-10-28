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
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

// saveRows2SQL save rows result into sql file
func saveRows2SQL(s *SaveStruct, rows *sql.Rows) error {

	file, err := os.Create(s.Config.File)
	if err != nil {
		return err
	}
	defer file.Close()

	// insert prefix
	insertPrefix, err := s.Config.SQLInsertPrefix(s.Config.DBParserColumnNames(s.Status.Header))
	if err != nil {
		return err
	}

	// header & columns
	headerColumns := s.Config.DBParseColumnTypes(s.Status.Header)
	columns := make([]interface{}, len(s.Status.Header))
	cols := make([]interface{}, len(s.Status.Header))
	for j := range columns {
		cols[j] = &columns[j]
	}

	// write every row into sql
	w := bufio.NewWriterSize(file, s.Config.MaxBufferSize)
	var sqlCounter int
	for rows.Next() {
		s.Status.Lines++
		// limit return rows
		if s.Config.Limit != 0 && s.Status.Lines > s.Config.Limit {
			s.Status.Lines = s.Config.Limit
			break
		}

		// scan columns
		if err := rows.Scan(cols...); err != nil {
			return err
		}

		// data mask
		values := make([]sql.NullString, len(columns))
		for j, col := range columns {
			if col == nil {
				values[j] = sql.NullString{String: s.Config.NULLString, Valid: false}
			} else {
				switch col.(type) {
				case []byte:
					values[j] = sql.NullString{String: string(col.([]byte)), Valid: true}
				case []string:
					values[j] = sql.NullString{String: s.Config.ParseArray(col.([]string)), Valid: true}
				default:
					values[j] = sql.NullString{String: fmt.Sprint(col), Valid: true}
				}
				// data mask
				valueMask, err := s.Masker.Mask(s.Status.Header[j].Name(), values[j].String)
				if err != nil {
					return err
				}
				values[j] = sql.NullString{String: valueMask, Valid: true}
			}
		}

		// concat values
		valuesStr, err := s.Config.SQLInsertValues(headerColumns, values)
		if err != nil {
			return err
		}

		// write sql
		sqlCounter++
		_, err = w.WriteString(s.Config.SQLMultiValues(sqlCounter, insertPrefix, valuesStr))
		if err != nil {
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	// last semicolon
	if s.Config.ExtendedInsert > 0 && s.Status.Lines > 0 && sqlCounter%s.Config.ExtendedInsert != 0 {
		_, err = w.WriteString(";\n")
		if err != nil {
			return err
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}
	return err
}
