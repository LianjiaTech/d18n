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

	"github.com/LianjiaTech/d18n/common"

	json "github.com/json-iterator/go"
)

func (d *DetectStruct) detectJSON() error {
	var err error

	fd, err := os.Open(d.Config.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	iter := json.Parse(json.ConfigDefault, fd, bufio.MaxScanTokenSize)
	switch iter.WhatIsNext() {
	case json.ArrayValue:
		if iter.WhatIsNext() == json.ArrayValue {
			if iter.ReadArrayCB(d.jsonDetectRow); iter.Error != nil {
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
func (d *DetectStruct) jsonDetectRow(iterator *json.Iterator) bool {
	var row []string
	if ok := iterator.ReadArrayCB(
		// callback function parse cell
		func(i *json.Iterator) bool {
			var elem string
			i.ReadVal(&elem)
			row = append(row, elem)
			return true
		}); ok {
		d.Status.Lines++

		// check column names
		if d.Status.Lines == 1 {
			if !d.Config.NoHeader && d.Config.Schema == "" {
				for _, r := range row {
					d.Status.Header = append(d.Status.Header, common.HeaderColumn{Name: r})
				}
			}
			d.checkHeader()
			if !d.Config.NoHeader {
				return true
			}
		}

		// skip lines
		if d.Status.Lines <= d.Config.SkipLines {
			return true
		}
		if d.Config.Limit > 0 &&
			(d.Status.Lines-d.Config.SkipLines) > d.Config.Limit {
			return false
		}

		// check value
		for j, value := range row {
			d.Status.Columns[d.Status.Header[j].Name] = append(d.Status.Columns[d.Status.Header[j].Name], d.checkValue(value)...)
		}
	}
	return true
}
