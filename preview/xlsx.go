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

package preview

import (
	"fmt"

	"d18n/common"

	"github.com/tealeg/xlsx/v3"
)

// PreviewXlsx ...
func previewXlsx(p *PreviewStruct) error {
	if p.Config.Preview == 0 {
		return nil
	}

	opts := xlsx.RowLimit(p.Config.Preview)
	wb, err := xlsx.OpenFile(p.Config.File, opts)
	if err != nil {
		return err
	}

	if len(wb.Sheets) > 0 {
		for i := 0; i < p.Config.Preview && i < wb.Sheets[0].MaxRow; i++ {
			row, err := wb.Sheets[0].Row(i)
			if err != nil {
				return err
			}
			for j := 0; j < wb.Sheets[0].MaxCol; j++ {
				fmt.Print(row.GetCell(j), "\t")
			}
			fmt.Println() // add line feed
		}
	}

	if p.Config.Verbose {
		watermark, err := common.GetXlsxWatermark(p.Config.File)
		if err == nil && watermark != "" {
			println("\nWatermark:", watermark)
		}
	}
	return nil
}
