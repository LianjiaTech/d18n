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

package detect

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/LianjiaTech/d18n/common"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func (d *DetectStruct) detectSQL() error {
	var err error
	f, err := os.Open(d.Config.File)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer([]byte{}, d.Config.MaxBufferSize)
	s.Split(common.SQLReadLine)

	// new sql parser
	p := parser.New()
	if d.Config.ANSIQuotes {
		mode, _ := mysql.GetSQLMode("ANSI_QUOTES")
		p.SetSQLMode(mode)
	}

	for s.Scan() {
		d.Status.Lines++

		if strings.TrimSpace(s.Text()) == "" {
			continue
		}

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

			// check column names
			if d.Status.Lines == 1 {
				if !d.Config.NoHeader && d.Config.Schema == "" {
					for _, col := range stmtNode.Columns {
						d.Status.Header = append(d.Status.Header, common.HeaderColumn{Name: col.String()})
					}
				}
				d.checkHeader()
			}

			// check value
			for _, r := range stmtNode.Lists {
				var row []string
				for _, cell := range r {
					var buf bytes.Buffer
					cell.Format(&buf)
					row = append(row, buf.String())
				}

				for j, value := range row {
					d.Status.Columns[d.Status.Header[j].Name] = append(d.Status.Columns[d.Status.Header[j].Name], d.checkValue(value)...)
				}
			}
		default:
			return fmt.Errorf(common.WrongSQLFormat)
		}
	}
	if s.Err() != nil {
		// bufio.ErrTooLong 1. raw data too large, 2. missing quotes or error comment
		return fmt.Errorf("line: %d, %s", d.Status.Lines+1, s.Err().Error())
	}
	return err
}
