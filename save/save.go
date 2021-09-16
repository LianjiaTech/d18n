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

package save

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"d18n/common"
)

type SaveStatus struct {
	Header        []*sql.ColumnType // Column Header, include column type info
	Lines         int               // lines save into file
	QueryTimeCost int64             // query time cost
	SaveTimeCost  int64             // save file time cost
	TimeCost      int64             // total time cost
}

var saveStatus SaveStatus

func Save() error {
	// execute sql and get all result rows
	queryStartTime := time.Now().UnixNano()
	rows, err := common.QueryRows()
	if err != nil {
		return err
	}
	defer rows.Close()

	queryEndTime := time.Now().UnixNano()

	// save rows result
	saveStartTime := time.Now().UnixNano()
	err = saveRows(rows)
	if err != nil {
		err = fmt.Errorf("line: %d, %s", saveStatus.Lines, err.Error())
	}
	saveEndTime := time.Now().UnixNano()

	// update time cost
	saveStatus.QueryTimeCost = queryEndTime - queryStartTime
	saveStatus.SaveTimeCost = saveEndTime - saveStartTime
	saveStatus.TimeCost = saveEndTime - queryStartTime

	return err
}

// saveRows ...
func saveRows(rows *sql.Rows) error {
	var err error

	// get column header
	saveStatus.Header, err = rows.ColumnTypes()
	if err != nil {
		return err
	}

	// file type switch
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(common.Cfg.File), "."))
	switch suffix {
	case "": // stdout ascii table
		if strings.EqualFold(common.Cfg.File, "stdout") {
			common.Cfg.Comma = '\t'
			err = saveRows2CSV(rows)
		} else {
			err = saveRows2ASCII(rows)
		}
	case "tsv": // tab-separated values
		common.Cfg.Comma = '\t'
		err = saveRows2CSV(rows)
	case "txt": // space-separated values
		common.Cfg.Comma = ' '
		err = saveRows2CSV(rows)
	case "psv": // pipe-separated values
		common.Cfg.Comma = '|'
		err = saveRows2CSV(rows)
	case "csv": // comma-separated values
		common.Cfg.Comma = ','
		err = saveRows2CSV(rows)
	case "html": // html
		err = saveRows2HTML(rows)
	case "xlsx": // microsoft office excel
		err = saveRows2XLSX(rows)
	case "sql": // sql file
		err = saveRows2SQL(rows)
	case "json": // json file, first element is column name, others are values
		err = saveRows2JSON(rows)
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}

	return err
}

// CheckStatus check SaveRows status at last
func CheckStatus() error {
	var err error

	// like QueryRow, should get result, but return empty set raise error
	if common.Cfg.CheckEmpty && saveStatus.Lines == 0 {
		err = fmt.Errorf(common.WrongEmptySet)
	}

	// verbose mode print
	if !common.Cfg.Verbose {
		return err
	}
	println(
		"Get rows:", saveStatus.Lines,
		"Query cost:", fmt.Sprint(time.Duration(saveStatus.QueryTimeCost)*time.Nanosecond),
		"Save cost:", fmt.Sprint(time.Duration(saveStatus.SaveTimeCost)*time.Nanosecond),
		"Total Cost:", fmt.Sprint(time.Duration(saveStatus.TimeCost)*time.Nanosecond),
	)
	return err
}
