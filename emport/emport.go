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
	"path/filepath"
	"strings"
	"time"

	"d18n/common"
)

var emportStatus EmportStatus

type EmportStatus struct {
	Header   []common.HeaderColumn // Column Header, include column type info
	Lines    int                   // lines save into file
	TimeCost int64                 // time cost
}

func Emport() error {
	emportStartTime := time.Now().UnixNano()
	// new *sql.DB, by pass error
	conn, _ := common.NewConnection()

	err := emportRows(conn)
	if err != nil {
		err = fmt.Errorf("line: %d, error: %s", emportStatus.Lines, err.Error())
	}

	emportEndTime := time.Now().UnixNano()
	emportStatus.TimeCost = emportEndTime - emportStartTime

	return err
}

func emportRows(conn *sql.DB) error {
	var err error
	// get Header
	emportStatus.Header, err = common.ParseSchema()
	if err != nil {
		return err
	}

	// exec sql before load data, e.g., foreign key, lock table write ...
	err = emportPrefixExec(conn)
	if err != nil {
		return err
	}

	// file type switch
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(common.Cfg.File), "."))
	switch suffix {
	case "tsv": // tab-separated values
		common.Cfg.Comma = '\t'
		err = emportCSV(conn)
	case "txt": // space-separated values
		common.Cfg.Comma = ' '
		err = emportCSV(conn)
	case "psv": // pipe-separated values
		common.Cfg.Comma = '|'
		err = emportCSV(conn)
	case "csv": // comma-separated values
		common.Cfg.Comma = ','
		err = emportCSV(conn)
	case "xlsx": // microsoft office excel
		err = emportXlsx(conn)
	case "html": // html format
		err = emportHTML(conn)
	case "sql": // sql file
		err = emportSQL(conn)
	case "json": // json file
		err = emportJSON(conn)
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}
	if err != nil {
		return err
	}

	// exec sql before load data, e.g., foreign key, lock table write ...
	err = emportSuffixExec(conn)

	return err
}

func emportPrefixExec(conn *sql.DB) error {
	var err error

	// change foreign key checks
	if common.Cfg.DisableForeignKeyChecks {
		err = common.SetForeignKeyChecks(!common.Cfg.DisableForeignKeyChecks, conn, common.Cfg.Table)
	}

	return err
}

func emportSuffixExec(conn *sql.DB) error {
	var err error

	// change foreign key checks
	if common.Cfg.DisableForeignKeyChecks {
		err = common.SetForeignKeyChecks(common.Cfg.DisableForeignKeyChecks, conn, common.Cfg.Table)
	}

	return err
}

// executeSQL ...
func executeSQL(sql string, conn *sql.DB) error {
	var err error

	if common.DBAvailable(conn) {
		_, err = conn.Exec(sql)
	} else {
		fmt.Print(sql)
	}

	return err
}

func CheckStatus() error {
	var err error

	if common.Cfg.SkipLines > 0 {
		println("Skip Lines:", common.Cfg.SkipLines)
	}

	// verbose mode print
	if !common.Cfg.Verbose {
		return err
	}
	println(
		"Import Rows:", emportStatus.Lines,
		"Total Cost:", fmt.Sprint(time.Duration(emportStatus.TimeCost)*time.Nanosecond),
	)

	return err
}
