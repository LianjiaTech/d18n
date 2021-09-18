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
	"bufio"
	"fmt"
	"os"

	"d18n/common"

	json "github.com/json-iterator/go"
)

func detectJSON() error {
	var err error

	fd, err := os.Open(common.Cfg.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	iter := json.Parse(json.ConfigDefault, fd, bufio.MaxScanTokenSize)
	switch iter.WhatIsNext() {
	case json.ArrayValue:
		if iter.WhatIsNext() == json.ArrayValue {
			if iter.ReadArrayCB(jsonDetectRow); iter.Error != nil {
				return iter.Error
			}
		} else {
			return fmt.Errorf(common.WrongJSONFormat)
		}
	default:
		return fmt.Errorf(common.WrongJSONFormat)
	}

	return err
}

// jsonDetectRow callback function of json-interator
func jsonDetectRow(iterator *json.Iterator) bool {
	var row []string
	if ok := iterator.ReadArrayCB(
		// callback function parse cell
		func(i *json.Iterator) bool {
			var elem string
			i.ReadVal(&elem)
			row = append(row, elem)
			return true
		}); ok {
		detectStatus.Lines++

		// check column names
		if detectStatus.Lines == 1 {
			if !common.Cfg.NoHeader && common.Cfg.Schema == "" {
				for _, r := range row {
					detectStatus.Header = append(detectStatus.Header, common.HeaderColumn{Name: r})
				}
			}
			checkFileHeader(detectStatus.Header)
			if !common.Cfg.NoHeader {
				return true
			}
		}

		// skip lines
		if detectStatus.Lines <= common.Cfg.SkipLines {
			return true
		}
		if common.Cfg.Limit > 0 &&
			(detectStatus.Lines-common.Cfg.SkipLines) > common.Cfg.Limit {
			return false
		}

		// check value
		for j, value := range row {
			detectStatus.Columns[detectStatus.Header[j].Name] = append(detectStatus.Columns[detectStatus.Header[j].Name], checkValue(value)...)
		}
	}
	return true
}
