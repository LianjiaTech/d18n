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

	json "github.com/json-iterator/go"
)

func emportJSON(e *EmportStruct, conn *sql.DB) error {
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
	iter := json.Parse(json.ConfigDefault, fd, bufio.MaxScanTokenSize)

	var sql string
	var sqlCounter int
	switch iter.WhatIsNext() {
	case json.ArrayValue:
		if iter.WhatIsNext() == json.ArrayValue {
			if iter.ReadArrayCB(
				// callback function stream read a line of data
				func(iterator *json.Iterator) bool {
					var row []string
					if ok := iterator.ReadArrayCB(
						// callback function parse cell. true continue, false break loop
						func(i *json.Iterator) bool {
							var elem string
							i.ReadVal(&elem)
							row = append(row, elem)
							return true
						}); ok {
						e.Status.Lines++

						// skip header line
						if e.Status.Lines == 1 && !common.Cfg.NoHeader {
							return true
						}

						// SkipLines
						if e.Status.Lines <= common.Cfg.SkipLines {
							return true
						}
						if common.Cfg.Limit > 0 &&
							(e.Status.Lines-common.Cfg.SkipLines) > common.Cfg.Limit {
							return false
						}

						// ignore blank lines
						if common.Cfg.IgnoreBlank && len(row) == 0 {
							return true
						}

						//  mask data
						row, err = e.Masker.MaskRow(e.Status.Header, row)
						if err != nil {
							iterator.Error = err
							return false
						}

						values, err := common.SQLInsertValues(e.Status.Header, common.DBParseNullString(e.Status.Header, row))
						if err != nil {
							iterator.Error = err
							return false
						}

						// extended-insert
						sqlCounter++
						sql += common.SQLMultiValues(sqlCounter, insertPrefix, values)
						if common.Cfg.ExtendedInsert <= 1 || sqlCounter%common.Cfg.ExtendedInsert == 0 {
							err = executeSQL(sql, conn)
							if err != nil {
								iterator.Error = err
								return false
							}
							sql = ""
						}
					}
					return true
				},
			); iter.Error != nil {
				return iter.Error
			}
		} else {
			return fmt.Errorf(common.WrongJSONFormat)
		}
	default:
		return fmt.Errorf(common.WrongJSONFormat)
	}

	// execute last SQL
	if sql != "" {
		err = executeSQL(sql, conn)
	}

	return err
}
