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
	"strings"

	"github.com/LianjiaTech/d18n/common"
)

// emportSQL import sql file into database
func emportSQL(e *EmportStruct, conn *sql.DB) error {
	var err error
	f, err := os.Open(e.Config.File)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer([]byte{}, e.Config.MaxBufferSize)
	s.Split(common.SQLReadLine)

	var sqlCounter int
	for s.Scan() {
		e.Status.Lines++

		// read one sql
		sqlString := strings.TrimSpace(s.Text())

		// SkipLines
		if e.Status.Lines <= e.Config.SkipLines {
			continue
		}
		if e.Config.Limit > 0 &&
			(e.Status.Lines-e.Config.SkipLines) > e.Config.Limit {
			e.Status.Lines = e.Config.Limit
			break
		}

		// ignore blank lines
		if e.Config.IgnoreBlank && strings.TrimSpace(sqlString) == "" {
			continue
		}

		// execute SQL
		sqlCounter++
		err = e.executeSQL(sqlString, conn)
		if err != nil {
			return err
		}
	}
	if s.Err() != nil {
		// bufio.ErrTooLong 1. raw data too large, 2. missing quotes or error comment
		return fmt.Errorf("line: %d, %s", e.Status.Lines+1, s.Err().Error())
	}

	e.Status.Rows = sqlCounter
	return err
}
