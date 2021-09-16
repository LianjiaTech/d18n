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
	"d18n/mask"

	xlsx "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func emportXlsx(conn *sql.DB) error {
	fd, err := xlsx.OpenFile(common.Cfg.File)
	if err != nil {
		return err
	}

	sheets := fd.GetSheetList()
	if len(sheets) > 0 {
		rows, err := fd.Rows(sheets[0])
		if err != nil {
			return err
		}

		insertPrefix, err := common.SQLInsertPrefix(common.DBParseHeaderColumn(emportStatus.Header))
		if err != nil {
			return err
		}

		var sql string
		var sqlCounter int
		for rows.Next() {
			emportStatus.Lines++
			// read row
			row, err := rows.Columns()
			if err != nil {
				return err
			}

			// skip header line
			if emportStatus.Lines == 1 && !common.Cfg.NoHeader {
				continue
			}

			// SkipLines
			if emportStatus.Lines <= common.Cfg.SkipLines {
				continue
			}
			if common.Cfg.Limit > 0 &&
				(emportStatus.Lines-common.Cfg.SkipLines) > common.Cfg.Limit {
				break
			}

			// ignore blank lines
			if common.Cfg.IgnoreBlank && len(row) == 0 {
				continue
			}

			// mask data
			if common.Cfg.IgnoreBlank {
				rowLen := len(row)
				// ignore extra blank cell
				if len(emportStatus.Header) < len(row) {
					rowLen = len(emportStatus.Header)
				}
				row, err = mask.MaskRow(emportStatus.Header, row[:rowLen])
			} else {
				row, err = mask.MaskRow(emportStatus.Header, row)
			}
			if err != nil {
				return err
			}

			// concat sql
			values, err := common.SQLInsertValues(emportStatus.Header, common.DBParseNullString(emportStatus.Header, row))
			if err != nil {
				return err
			}

			// extended-insert
			sqlCounter++
			sql += common.SQLMultiValues(sqlCounter, insertPrefix, values)
			if common.Cfg.ExtendedInsert <= 1 || sqlCounter%common.Cfg.ExtendedInsert == 0 {
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

	} else {
		return fmt.Errorf("empty xlsx file")
	}
	return err
}
