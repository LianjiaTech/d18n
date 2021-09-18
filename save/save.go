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
	"d18n/mask"
)

type saveStatus struct {
	Header        []*sql.ColumnType // Column Header, include column type info
	Lines         int               // lines save into file
	QueryTimeCost int64             // query time cost
	SaveTimeCost  int64             // save file time cost
	TimeCost      int64             // total time cost
}

type SaveStruct struct {
	CommonConfig common.Config    // common config
	Status       saveStatus       // save status
	Masker       *mask.MaskStruct // masker
}

func NewSaveStruct(c common.Config) (*SaveStruct, error) {
	var s *SaveStruct
	m, err := mask.NewMaskStruct(c)
	if err != nil {
		return s, err
	}

	s = &SaveStruct{
		CommonConfig: c,
		Masker:       m,
	}
	return s, err
}

func (s *SaveStruct) Save() error {
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
	err = saveRows(s, rows)
	if err != nil {
		err = fmt.Errorf("line: %d, %s", s.Status.Lines, err.Error())
	}
	saveEndTime := time.Now().UnixNano()

	// update time cost
	s.Status.QueryTimeCost = queryEndTime - queryStartTime
	s.Status.SaveTimeCost = saveEndTime - saveStartTime
	s.Status.TimeCost = saveEndTime - queryStartTime

	return err
}

// saveRows ...
func saveRows(s *SaveStruct, rows *sql.Rows) error {
	var err error

	// get column header
	s.Status.Header, err = rows.ColumnTypes()
	if err != nil {
		return err
	}

	// file type switch
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(s.CommonConfig.File), "."))
	switch suffix {
	case "": // stdout ascii table
		if strings.EqualFold(s.CommonConfig.File, "stdout") {
			s.CommonConfig.Comma = '\t'
			err = saveRows2CSV(s, rows)
		} else {
			err = saveRows2ASCII(s, rows)
		}
	case "tsv": // tab-separated values
		s.CommonConfig.Comma = '\t'
		err = saveRows2CSV(s, rows)
	case "txt": // space-separated values
		s.CommonConfig.Comma = ' '
		err = saveRows2CSV(s, rows)
	case "psv": // pipe-separated values
		s.CommonConfig.Comma = '|'
		err = saveRows2CSV(s, rows)
	case "csv": // comma-separated values
		s.CommonConfig.Comma = ','
		err = saveRows2CSV(s, rows)
	case "html": // html
		err = saveRows2HTML(s, rows)
	case "xlsx": // microsoft office excel
		err = saveRows2XLSX(s, rows)
	case "sql": // sql file
		err = saveRows2SQL(s, rows)
	case "json": // json file, first element is column name, others are values
		err = saveRows2JSON(s, rows)
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}

	return err
}

// CheckStatus check SaveRows status at last
func (s *SaveStruct) CheckStatus() error {
	var err error

	// like QueryRow, should get result, but return empty set raise error
	if s.CommonConfig.CheckEmpty && s.Status.Lines == 0 {
		err = fmt.Errorf(common.WrongEmptySet)
	}

	// verbose mode print
	if !s.CommonConfig.Verbose {
		return err
	}
	println(
		"Get rows:", s.Status.Lines,
		"Query cost:", fmt.Sprint(time.Duration(s.Status.QueryTimeCost)*time.Nanosecond),
		"Save cost:", fmt.Sprint(time.Duration(s.Status.SaveTimeCost)*time.Nanosecond),
		"Total Cost:", fmt.Sprint(time.Duration(s.Status.TimeCost)*time.Nanosecond),
	)
	return err
}
