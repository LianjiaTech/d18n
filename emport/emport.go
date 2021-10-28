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

	"github.com/LianjiaTech/d18n/common"
	"github.com/LianjiaTech/d18n/mask"
)

type emportStatus struct {
	Header   []common.HeaderColumn // Column Header, include column type info
	Lines    int                   // file lines
	Rows     int                   // rows of values
	TimeCost int64                 // time cost
}

type EmportStruct struct {
	Status emportStatus     // emport status
	Masker *mask.MaskStruct // masker
	Config common.Config    //
}

func NewEmportStruct(c common.Config) (*EmportStruct, error) {
	var e *EmportStruct
	m, err := mask.NewMaskStruct(c.Mask)
	if err != nil {
		return e, err
	}

	e = &EmportStruct{
		Masker: m,
		Config: c,
		Status: emportStatus{
			Lines: 0,
		},
	}
	return e, nil
}

func (e *EmportStruct) Emport() error {
	emportStartTime := time.Now().UnixNano()
	// new *sql.DB, by pass error
	conn, _ := e.Config.NewConnection()

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
	e.Status.Header, err = e.Config.ParseSchema()
	if err != nil {
		return err
	}

	// exec sql before load data, e.g., foreign key, lock table write ...
	err = emportPrefixExec(e, conn)
	if err != nil {
		return err
	}

	// file type switch
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(e.Config.File), "."))
	switch suffix {
	case "tsv": // tab-separated values
		e.Config.Comma = '\t'
		err = emportCSV(e, conn)
	case "txt": // space-separated values
		e.Config.Comma = ' '
		err = emportCSV(e, conn)
	case "psv": // pipe-separated values
		e.Config.Comma = '|'
		err = emportCSV(e, conn)
	case "csv": // comma-separated values
		e.Config.Comma = ','
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
	if e.Config.DisableForeignKeyChecks {
		err = e.Config.SetForeignKeyChecks(!e.Config.DisableForeignKeyChecks, conn, e.Config.Table)
	}

	return err
}

func emportSuffixExec(e *EmportStruct, conn *sql.DB) error {
	var err error

	// change foreign key checks
	if e.Config.DisableForeignKeyChecks {
		err = e.Config.SetForeignKeyChecks(e.Config.DisableForeignKeyChecks, conn, e.Config.Table)
	}

	return err
}

// executeSQL ...
func (e *EmportStruct) executeSQL(sql string, conn *sql.DB) error {
	var err error

	if e.Config.DBAvailable(conn) {
		_, err = conn.Exec(sql)
	} else {
		fmt.Print(sql)
	}

	return err
}

func (e *EmportStruct) ShowStatus() error {
	var err error

	if e.Config.SkipLines > 0 {
		println("Skip Lines:", e.Config.SkipLines)
	}

	// verbose mode print
	if len(e.Config.Verbose) == 0 {
		return err
	}
	println(
		"File Lines:", e.Status.Lines,
		"Import Rows:", e.Status.Rows,
		"Total Cost:", fmt.Sprint(time.Duration(e.Status.TimeCost)*time.Nanosecond),
	)

	return err
}
