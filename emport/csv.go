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
	"encoding/csv"
	"io"
	"os"

	"d18n/common"
)

func emportCSV(e *EmportStruct, conn *sql.DB) error {
	var err error

	fd, err := os.Open(common.Cfg.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	insertPrefix, err := common.SQLInsertPrefix(common.DBParseHeaderColumn(e.Status.Header))
	if err != nil {
		return err
	}

	r := csv.NewReader(fd)
	r.Comma = common.Cfg.Comma
	var sql string
	var sqlCounter int
	for {
		e.Status.Lines++

		row, err := r.Read()
		if err == io.EOF { // end of file
			break
		} else if err != nil {
			return err
		}

		// skip header line
		if e.Status.Lines == 1 && !common.Cfg.NoHeader {
			continue
		}

		// SkipLines
		if e.Status.Lines <= common.Cfg.SkipLines {
			continue
		}
		if common.Cfg.Limit > 0 &&
			(e.Status.Lines-common.Cfg.SkipLines) > common.Cfg.Limit {
			break
		}

		// ignore blank lines
		if common.Cfg.IgnoreBlank && len(row) == 0 {
			continue
		}

		//  mask data
		row, err = e.Masker.MaskRow(e.Status.Header, row)
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
	}

	return err
}
