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
	"bufio"
	"database/sql"
	"fmt"
	"os"

	"d18n/common"

	"golang.org/x/net/html"
)

func emportHTML(e *EmportStruct, conn *sql.DB) error {
	var err error

	fd, err := os.Open(e.CommonConfig.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	insertPrefix, err := common.SQLInsertPrefix(common.DBParseHeaderColumn(e.Status.Header))
	if err != nil {
		return err
	}

	r := bufio.NewReaderSize(fd, e.CommonConfig.MaxBufferSize)
	token := html.NewTokenizer(r)

	var row []string
	var sql string
	var sqlCounter int
	for {

		t := token.Next()
		if t == html.ErrorToken {
			break
		}

		tag, _ := token.TagName()
		switch t {
		case html.StartTagToken:
			switch string(tag) {
			case "th":
				// skip header line
			case "td":
				token.Next()
				row = append(row, html.UnescapeString(string(token.Raw())))
			case "tr":
				e.Status.Lines++
			}
		case html.EndTagToken:
			switch string(tag) {
			case "tr":
				// skip header line
				if e.Status.Lines == 1 && !e.CommonConfig.NoHeader {
					continue
				}

				// ignore blank lines
				if e.CommonConfig.IgnoreBlank && len(row) == 0 {
					continue
				}

				if len(e.Status.Header) != len(row) {
					return fmt.Errorf(common.WrongColumnsCnt)
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
				if e.CommonConfig.ExtendedInsert <= 1 || sqlCounter%e.CommonConfig.ExtendedInsert == 0 {
					err = executeSQL(sql, conn)
					if err != nil {
						return err
					}
					sql = ""
				}

				// truncate row after new line
				row = []string{}
			}
		}

		// SkipLines
		if e.Status.Lines <= e.CommonConfig.SkipLines {
			continue
		}
		if e.CommonConfig.Limit > 0 &&
			(e.Status.Lines-e.CommonConfig.SkipLines) > e.CommonConfig.Limit {
			break
		}
	}

	// execute last SQL
	if sql != "" {
		err = executeSQL(sql, conn)
	}

	return err
}
