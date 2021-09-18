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

func detectCSV() error {

	var err error

	fd, err := os.Open(common.Cfg.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	r := csv.NewReader(fd)
	r.Comma = common.Cfg.Comma
	for {
		detectStatus.Lines++

		// read row
		row, err := r.Read()
		if err == io.EOF { // end of file
			break
		} else if err != nil {
			return err
		}

		// check column names
		if detectStatus.Lines == 1 {
			if !common.Cfg.NoHeader && common.Cfg.Schema == "" {
				for _, r := range row {
					detectStatus.Header = append(detectStatus.Header, common.HeaderColumn{Name: r})
				}
			}
			checkFileHeader(detectStatus.Header)
			if !common.Cfg.NoHeader {
				continue
			}
		}

		// SkipLines
		if detectStatus.Lines <= common.Cfg.SkipLines {
			continue
		}
		if common.Cfg.Limit > 0 &&
			(detectStatus.Lines-common.Cfg.SkipLines) > common.Cfg.Limit {
			break
		}

		// check value
		for j, value := range row {
			detectStatus.Columns[detectStatus.Header[j].Name] = append(detectStatus.Columns[detectStatus.Header[j].Name], checkValue(value)...)
		}
	}

	return err
}
