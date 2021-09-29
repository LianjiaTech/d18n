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
	"d18n/mask"
)

type emportStatus struct {
	Header   []common.HeaderColumn // Column Header, include column type info
	Lines    int                   // file lines
	Rows     int                   // rows of values
	TimeCost int64                 // time cost
}

type EmportStruct struct {
	Status       emportStatus     // emport status
	Masker       *mask.MaskStruct // masker
	CommonConfig common.Config    //
}

func NewEmportStruct(c common.Config) (*EmportStruct, error) {
	var e *EmportStruct
	m, err := mask.NewMaskStruct(c.Mask)
	if err != nil {
		return e, err
	}

	e = &EmportStruct{
		Masker:       m,
		CommonConfig: c,
	}
	return e, nil
}

func (e *EmportStruct) Emport() error {
	emportStartTime := time.Now().UnixNano()
	// new *sql.DB, by pass error
	conn, _ := common.NewConnection()

	err := emportRows(e, conn)
	if err != nil {
		err = fmt.Errorf("line: %d, error: %s", e.Status.Lines, err.Error())
	}

	emportEndTime := time.Now().UnixNano()
	e.Status.TimeCost = emportEndTime - emportStartTime

	return err
}

func emportRows(e *EmportStruct, conn *sql.DB) error {
	var err error
	// get Header
	e.Status.Header, err = common.ParseSchema()
	if err != nil {
		return err
	}

	// exec sql before load data, e.g., foreign key, lock table write ...
	err = emportPrefixExec(e, conn)
	if err != nil {
		return err
	}

	// file type switch
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(e.CommonConfig.File), "."))
	switch suffix {
	case "tsv": // tab-separated values
		e.CommonConfig.Comma = '\t'
		err = emportCSV(e, conn)
	case "txt": // space-separated values
		e.CommonConfig.Comma = ' '
		err = emportCSV(e, conn)
	case "psv": // pipe-separated values
		e.CommonConfig.Comma = '|'
		err = emportCSV(e, conn)
	case "csv": // comma-separated values
		e.CommonConfig.Comma = ','
		err = emportCSV(e, conn)
	case "xlsx": // microsoft office excel
		err = emportXlsx(e, conn)
	case "html": // html format
		err = emportHTML(e, conn)
	case "sql": // sql file
		err = emportSQL(e, conn)
	case "json": // json file
		err = emportJSON(e, conn)
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}
	if err != nil {
		return err
	}

	// exec sql before load data, e.g., foreign key, lock table write ...
	err = emportSuffixExec(e, conn)

	return err
}

func emportPrefixExec(e *EmportStruct, conn *sql.DB) error {
	var err error

	// change foreign key checks
	if e.CommonConfig.DisableForeignKeyChecks {
		err = common.SetForeignKeyChecks(!e.CommonConfig.DisableForeignKeyChecks, conn, e.CommonConfig.Table)
	}

	return err
}

func emportSuffixExec(e *EmportStruct, conn *sql.DB) error {
	var err error

	// change foreign key checks
	if e.CommonConfig.DisableForeignKeyChecks {
		err = common.SetForeignKeyChecks(e.CommonConfig.DisableForeignKeyChecks, conn, e.CommonConfig.Table)
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

func (e *EmportStruct) ShowStatus() error {
	var err error

	if e.CommonConfig.SkipLines > 0 {
		println("Skip Lines:", e.CommonConfig.SkipLines)
	}

	// verbose mode print
	if !e.CommonConfig.Verbose {
		return err
	}
	println(
		"File Lines:", e.Status.Lines,
		"Import Rows:", e.Status.Rows,
		"Total Cost:", fmt.Sprint(time.Duration(e.Status.TimeCost)*time.Nanosecond),
	)

	return err
}
