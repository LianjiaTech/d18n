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

package detect

import (
	"encoding/csv"
	"io"
	"os"

	"d18n/common"
)

func (d *DetectStruct) detectCSV() error {

	var err error

	fd, err := os.Open(d.Config.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	r := csv.NewReader(fd)
	r.Comma = d.Config.Comma
	for {
		d.Status.Lines++

		// read row
		row, err := r.Read()
		if err == io.EOF { // end of file
			break
		} else if err != nil {
			return err
		}

		// check column names
		if d.Status.Lines == 1 {
			if !d.Config.NoHeader && d.Config.Schema == "" {
				for _, r := range row {
					d.Status.Header = append(d.Status.Header, common.HeaderColumn{Name: r})
				}
			}
			d.checkHeader()
			if !d.Config.NoHeader {
				continue
			}
		}

		// SkipLines
		if d.Status.Lines <= d.Config.SkipLines {
			continue
		}
		if d.Config.Limit > 0 &&
			(d.Status.Lines-d.Config.SkipLines) > d.Config.Limit {
			break
		}

		// check value
		for j, value := range row {
			d.Status.Columns[d.Status.Header[j].Name] = append(d.Status.Columns[d.Status.Header[j].Name], d.checkValue(value)...)
		}
	}

	return err
}
