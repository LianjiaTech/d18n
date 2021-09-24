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

package emport

import (
	"database/sql"
	"fmt"

	"d18n/common"

	xlsx "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func emportXlsx(e *EmportStruct, conn *sql.DB) error {
	fd, err := xlsx.OpenFile(e.CommonConfig.File)
	if err != nil {
		return err
	}

	sheets := fd.GetSheetList()
	if len(sheets) > 0 {
		rows, err := fd.Rows(sheets[0])
		if err != nil {
			return err
		}

		insertPrefix, err := common.SQLInsertPrefix(common.DBParseHeaderColumn(e.Status.Header))
		if err != nil {
			return err
		}

		var sql string
		var sqlCounter int
		for rows.Next() {
			e.Status.Lines++
			// read row
			row, err := rows.Columns()
			if err != nil {
				return err
			}

			// skip header line
			if e.Status.Lines == 1 && !e.CommonConfig.NoHeader {
				continue
			}

			// SkipLines
			if e.Status.Lines <= e.CommonConfig.SkipLines {
				continue
			}
			if e.CommonConfig.Limit > 0 &&
				(e.Status.Lines-e.CommonConfig.SkipLines) > e.CommonConfig.Limit {
				break
			}

			// ignore blank lines
			if e.CommonConfig.IgnoreBlank && len(row) == 0 {
				continue
			}

			// mask data
			if e.CommonConfig.IgnoreBlank {
				rowLen := len(row)
				// ignore extra blank cell
				if len(e.Status.Header) < len(row) {
					rowLen = len(e.Status.Header)
				}
				row, err = e.Masker.MaskRow(e.Status.Header, row[:rowLen])
			} else {
				row, err = e.Masker.MaskRow(e.Status.Header, row)
			}
			if err != nil {
				return err
			}

			// concat sql
			values, err := common.SQLInsertValues(e.Status.Header, common.DBParseNullString(e.Status.Header, row))
			if err != nil {
				return err
			}

			// extended-insert
			sqlCounter++
			sql += common.SQLMultiValues(sqlCounter, insertPrefix, values)
			if e.CommonConfig.ExtendedInsert <= 1 || sqlCounter%e.CommonConfig.ExtendedInsert == 0 {
				err = executeSQL(sql, conn)
				if err != nil {
					return err
				}
				sql = ""
			}
		}

		// execute last SQL
		if sql != "" {
			err = executeSQL(sql, conn)
			if err != nil {
				return err
			}
		}
		e.Status.Rows = sqlCounter

	} else {
		return fmt.Errorf("empty xlsx file")
	}
	return err
}
