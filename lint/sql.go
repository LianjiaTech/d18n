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

package lint

import (
	"bufio"
	"bytes"
	"d18n/common"
	"fmt"
	"os"
	"strings"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func lintSQL() error {
	var err error
	f, err := os.Open(common.Cfg.File)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer([]byte{}, common.Cfg.MaxBufferSize)
	s.Split(common.SQLReadLine)

	// new sql parser
	p := parser.New()
	if common.Cfg.ANSIQuotes {
		mode, _ := mysql.GetSQLMode("ANSI_QUOTES")
		p.SetSQLMode(mode)
	}

	for s.Scan() {
		lintStatus.RowCount++

		stmt, err := p.ParseOneStmt(s.Text(), mysql.DefaultCharset, mysql.DefaultCollationName)
		if err != nil {
			return err
		}

		switch stmtNode := stmt.(type) {
		case *ast.InsertStmt: // INSERT, REPLACE
			// not support insert select clause
			if stmtNode.Select != nil {
				return fmt.Errorf(common.WrongSQLFormat)
			}

			// get header, table name & columns list
			if lintStatus.RowCount == 1 && !common.Cfg.NoHeader {
				for _, col := range stmtNode.Columns {
					lintStatus.Header = append(lintStatus.Header, col.String())
				}

				if common.Cfg.Table == "" {
					switch node := stmtNode.Table.TableRefs.Left.(type) {
					case *ast.TableSource:
						if n, ok := node.Source.(*ast.TableName); ok {
							common.Cfg.Table = n.Name.O
						}
					}
				}
			}

			// cell validation
			for _, r := range stmtNode.Lists {
				var row []string
				for _, cell := range r {
					var buf bytes.Buffer
					cell.Format(&buf)
					if common.Cfg.ANSIQuotes && strings.HasPrefix(buf.String(), "`") {
						return fmt.Errorf(common.WrongQuotesValue)
					}
					row = append(row, buf.String())
				}

				err = lintCell(lintStatus.RowCount, row)
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf(common.WrongSQLFormat)
		}
	}
	if s.Err() != nil {
		// bufio.ErrTooLong 1. raw data too large, 2. missing quotes or error comment
		return fmt.Errorf("line: %d, %s", lintStatus.CellCount+1, s.Err().Error())
	}
	return err
}
