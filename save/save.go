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
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/LianjiaTech/d18n/common"
	"github.com/LianjiaTech/d18n/mask"
	ora "github.com/sijms/go-ora/v2"
)

type saveStatus struct {
	Header        []*sql.ColumnType // Column Header, include column type info
	Lines         int               // lines save into file
	QueryTimeCost int64             // query time cost
	SaveTimeCost  int64             // save file time cost
	TimeCost      int64             // total time cost
}

type SaveStruct struct {
	Config common.Config    // common config
	Status saveStatus       // save status
	Masker *mask.MaskStruct // masker
	Rand   *rand.Rand
}

func NewSaveStruct(c common.Config) (*SaveStruct, error) {
	var s *SaveStruct
	m, err := mask.NewMaskStruct(c.Mask)
	if err != nil {
		return s, err
	}

	s = &SaveStruct{
		Config: c,
		Masker: m,
		Rand:   rand.New(rand.NewSource(c.RandSeed)),
	}
	return s, err
}

func (s *SaveStruct) FieldName(i int) string {
	name := s.Status.Header[i].Name()
	if n, ok := s.Config.FieldsAliasMap[name]; ok {
		name = n
	}
	return name
}

func (s *SaveStruct) Save() error {
	// execute sql and get all result rows
	queryStartTime := time.Now().UnixNano()
	rows, err := s.Config.QueryRows()
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
	ext := strings.ToLower(strings.Trim(filepath.Ext(s.Config.File), "."))
	if ext == "" {
		ext = strings.ToLower(strings.Trim(filepath.Base(s.Config.File), "."))
		s.Config.File = "stdout"
	}
	switch ext {
	case "", "stdout": // stdout ascii table
		err = saveRows2ASCII(s, rows)
	case "tsv": // tab-separated values
		s.Config.Comma = '\t'
		err = saveRows2CSV(s, rows)
	case "txt": // space-separated values
		s.Config.Comma = ' '
		err = saveRows2CSV(s, rows)
	case "psv": // pipe-separated values
		s.Config.Comma = '|'
		err = saveRows2CSV(s, rows)
	case "csv": // comma-separated values
		s.Config.Comma = ','
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
		err = fmt.Errorf("not support extension: " + ext)
	}

	return err
}

// sample ...
func (s *SaveStruct) sample() bool {
	var is bool
	if s.Config.Sample >= 1 || s.Config.Sample <= 0 || // wrong config
		s.Rand.Float64() <= s.Config.Sample {
		is = true
	}
	return is
}

// ShowStatus check SaveRows status at last
func (s *SaveStruct) ShowStatus() error {
	var err error

	// like QueryRow, should get result, but return empty set raise error
	if s.Config.CheckEmpty && s.Status.Lines == 0 {
		err = fmt.Errorf(common.WrongEmptySet)
	}

	// verbose mode print
	if len(s.Config.Verbose) == 0 {
		return err
	}
	if len(s.Config.Verbose) >= 1 {
		println(
			"Get rows:", s.Status.Lines,
			"Query cost:", fmt.Sprint(time.Duration(s.Status.QueryTimeCost)*time.Nanosecond),
		)
	}
	if len(s.Config.Verbose) >= 2 {
		println(
			"Save cost:", fmt.Sprint(time.Duration(s.Status.SaveTimeCost)*time.Nanosecond),
			"Total Cost:", fmt.Sprint(time.Duration(s.Status.TimeCost)*time.Nanosecond),
		)
	}
	return err
}

func (s *SaveStruct) TimeFormat(t time.Time, format string) string {
	if format == "" {
		format = common.DATETIME_FORMAT
	}
	return time.Time(t).Format(format)
}

func (s *SaveStruct) String(col interface{}, ty *sql.ColumnType) string {
	var str string
	// fmt.Printf("Type: %T, Value: %v, Ty: %s\n", col, col, ty.DatabaseTypeName())
	switch col.(type) {
	case []byte:
		var special bool
		switch s.Config.Server {
		case "mssql", "sqlserver":
			switch ty.DatabaseTypeName() {
			case "UNIQUEIDENTIFIER":
				if u := col.([]byte); len(u) == 16 {
					str = fmt.Sprintf("%X-%X-%X-%X-%X", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
					special = true
				}
			}
		case "oracle":
			switch ty.DatabaseTypeName() {
			case "RAW", "LongRaw", "VarRaw", // RAW
				"OCIBlobLocator": // BLOB
				str = fmt.Sprintf("%X", col.([]byte))
				special = true
			}
		}
		if !special {
			str = string(col.([]byte))
		}
	case []string:
		str = s.Config.ParseArray(col.([]string))
	case float32: // oracle number
		str = strconv.FormatFloat(float64(col.(float32)), 'f', -1, 64)
	case float64: // oracle number
		str = strconv.FormatFloat(col.(float64), 'f', -1, 64)
	case time.Time: // oracle date/datetime
		str = s.TimeFormat(col.(time.Time), common.DATETIME_FORMAT)
	case ora.TimeStamp: // oracle timestamp
		str = s.TimeFormat(time.Time(col.(ora.TimeStamp)), common.DATETIME_FORMAT)
	default:
		switch s.Config.Server {
		case "oracle":
			switch ty.DatabaseTypeName() {
			case "CHAR": // Oracle will auto fill space with CHAR type
				str = strings.TrimSpace(fmt.Sprint(col))
			default:
				str = fmt.Sprint(col)
			}
		case "sqlserver", "mssql":
			switch ty.DatabaseTypeName() {
			case "CHAR": // SQL Server will auto fill space with CHAR type
				str = strings.TrimSpace(fmt.Sprint(col))
			default:
				str = fmt.Sprint(col)
			}
		case "presto", "trino":
			switch ty.DatabaseTypeName() {
			case "CHAR": // SQL Server will auto fill space with CHAR type
				str = strings.TrimSpace(fmt.Sprint(col))
			default:
				str = fmt.Sprint(col)
			}
		default:
			str = fmt.Sprint(col)
		}
	}
	return str
}
