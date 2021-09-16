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

	"d18n/common"
	"d18n/mask"

	"github.com/tealeg/xlsx/v3"
)

// Excel limits
const (
	ExcelMaxRows      = 1048576 // 1M
	ExcelMaxColumns   = 16384   // 16K
	ExcelMaxCellChars = 32767   // 32K
)

// saveRows2XLSX save rows result into xlsx format file
func saveRows2XLSX(rows *sql.Rows) error {
	file := xlsx.NewFile()
	// create new sheet
	sheet, err := file.AddSheet("result")
	if err != nil {
		return err
	}

	// check columns count
	if len(saveStatus.Header) > ExcelMaxColumns {
		return fmt.Errorf("excel max columns(%d) exceeded", ExcelMaxColumns)
	}

	// set table header with column name
	if !common.Cfg.NoHeader {
		sheetHeader := sheet.AddRow()
		sheetHeader.SetHeight(12.5) // https://github.com/tealeg/xlsx/issues/647
		for _, header := range saveStatus.Header {
			cell := sheetHeader.AddCell()
			cell.Value = header.Name()
		}
	}

	// init columns
	columns := make([]interface{}, len(saveStatus.Header))
	cols := make([]interface{}, len(saveStatus.Header))
	for j := range columns {
		cols[j] = &columns[j]
	}

	var bufSize int
	for rows.Next() {
		saveStatus.Lines++

		// Check excel limit
		if saveStatus.Lines > ExcelMaxRows {
			return fmt.Errorf("excel max rows(%d) exceeded", ExcelMaxRows)
		}
		if bufSize > common.Cfg.ExcelMaxFileSize {
			return fmt.Errorf("excel max file size(%d) exceeded", common.Cfg.ExcelMaxFileSize)
		}

		// limit return rows
		if common.Cfg.Limit != 0 && saveStatus.Lines > common.Cfg.Limit {
			break
		}

		// scan columns
		if err := rows.Scan(cols...); err != nil {
			return err
		}

		values := make([]string, len(columns))
		for j, col := range columns {
			if col == nil {
				values[j] = common.Cfg.NULLString
			} else {
				switch col.(type) {
				case []byte:
					values[j] = string(col.([]byte))
				case []string:
					values[j] = common.ParseArray(col.([]string))
				default:
					values[j] = fmt.Sprint(col)
				}

				// data mask
				values[j], err = mask.Mask(saveStatus.Header[j].Name(), values[j])
				if err != nil {
					return err
				}

				// hex-blob
				values[j], _ = common.HexBLOB(saveStatus.Header[j].Name(), values[j])
			}
		}

		sheetRow := sheet.AddRow()
		sheetRow.SetHeight(12.5) // https://github.com/tealeg/xlsx/issues/647
		for _, v := range values {
			cell := sheetRow.AddCell()
			if len(v) > ExcelMaxCellChars {
				return fmt.Errorf("excel max cell characters(%d) exceeded", ExcelMaxCellChars)
			}
			cell.Value = v
			bufSize += len(cell.Value)
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	// save to file
	err = file.Save(common.Cfg.File)
	if err != nil {
		return err
	}

	if common.Cfg.Watermark != "" {
		err = common.SetXlsxWatermark(common.Cfg.File, common.Cfg.Watermark)
	}

	return err
}
