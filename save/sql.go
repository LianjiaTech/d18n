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

	"d18n/common"
	"d18n/mask"
)

// saveRows2SQL save rows result into sql file
func saveRows2SQL(rows *sql.Rows) error {

	file, err := os.Create(common.Cfg.File)
	if err != nil {
		return err
	}
	defer file.Close()

	// insert prefix
	insertPrefix, err := common.SQLInsertPrefix(common.DBParserColumnNames(saveStatus.Header))
	if err != nil {
		return err
	}

	// header & columns
	headerColumns := common.DBParseColumnTypes(saveStatus.Header)
	columns := make([]interface{}, len(saveStatus.Header))
	cols := make([]interface{}, len(saveStatus.Header))
	for j := range columns {
		cols[j] = &columns[j]
	}

	// write every row into sql
	w := bufio.NewWriterSize(file, common.Cfg.MaxBufferSize)
	var sqlCounter int
	for rows.Next() {
		saveStatus.Lines++
		// limit return rows
		if common.Cfg.Limit != 0 && saveStatus.Lines > common.Cfg.Limit {
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
				values[j] = sql.NullString{String: common.Cfg.NULLString, Valid: false}
			} else {
				switch col.(type) {
				case []byte:
					values[j] = sql.NullString{String: string(col.([]byte)), Valid: true}
				case []string:
					values[j] = sql.NullString{String: common.ParseArray(col.([]string)), Valid: true}
				default:
					values[j] = sql.NullString{String: fmt.Sprint(col), Valid: true}
				}
				// data mask
				valueMask, err := mask.Mask(saveStatus.Header[j].Name(), values[j].String)
				if err != nil {
					return err
				}
				values[j] = sql.NullString{String: valueMask, Valid: true}
			}
		}

		// concat values
		valuesStr, err := common.SQLInsertValues(headerColumns, values)
		if err != nil {
			return err
		}

		// write sql
		sqlCounter++
		_, err = w.WriteString(common.SQLMultiValues(sqlCounter, insertPrefix, valuesStr))
		if err != nil {
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	// last semicolon
	if common.Cfg.ExtendedInsert > 0 && saveStatus.Lines > 0 && sqlCounter%common.Cfg.ExtendedInsert != 0 {
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
