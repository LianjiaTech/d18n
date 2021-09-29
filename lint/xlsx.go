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

package lint

import (
	"fmt"

	"d18n/common"

	xlsx "github.com/360EntSecGroup-Skylar/excelize/v2"
)

// lintXlsx lint excel file
func (l *LintStruct) lintXlsx() error {
	f, err := xlsx.OpenFile(l.CommonConfig.File)
	if err != nil {
		return err
	}

	sheets := f.GetSheetList()
	if len(sheets) > 0 {
		rows, err := f.Rows(sheets[0])
		if err != nil {
			return err
		}

		for rows.Next() {
			// read row
			row, err := rows.Columns()
			if err != nil {
				return err
			}
			l.Status.RowCount++

			// add header
			if l.Status.RowCount == 1 && !l.CommonConfig.NoHeader {
				l.Status.Header = row
			}

			// line validation, check empty line
			if len(row) == 0 && !l.CommonConfig.IgnoreBlank {
				return fmt.Errorf(common.WrongColumnsCnt)
			}

			// cell validation
			err = l.lintCell(l.Status.RowCount, row)
			if err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("empty xlsx file")
	}
	return nil
}
