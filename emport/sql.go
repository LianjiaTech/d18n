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

	"d18n/common"
)

// emportSQL import sql file into database
func emportSQL(conn *sql.DB) error {
	var err error
	f, err := os.Open(common.Cfg.File)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer([]byte{}, common.Cfg.MaxBufferSize)
	s.Split(common.SQLReadLine)

	for s.Scan() {
		emportStatus.Lines++

		// read one sql
		sqlString := strings.TrimSpace(s.Text())

		// SkipLines
		if emportStatus.Lines <= common.Cfg.SkipLines {
			continue
		}
		if common.Cfg.Limit > 0 &&
			(emportStatus.Lines-common.Cfg.SkipLines) > common.Cfg.Limit {
			break
		}

		// ignore blank lines
		if common.Cfg.IgnoreBlank && strings.TrimSpace(sqlString) == "" {
			continue
		}

		// execute SQL
		err = executeSQL(sqlString, conn)
		if err != nil {
			return err
		}
	}
	if s.Err() != nil {
		// bufio.ErrTooLong 1. raw data too large, 2. missing quotes or error comment
		return fmt.Errorf("line: %d, %s", emportStatus.Lines+1, s.Err().Error())
	}
	return err
}
