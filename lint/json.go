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
	"bufio"
	"fmt"
	"os"

	"github.com/LianjiaTech/d18n/common"

	json "github.com/json-iterator/go"
)

// jsonLintRow callback function of json-interator
func (l *LintStruct) jsonLintRow(iter *json.Iterator) bool {
	var err error
	var row []string
	if ok := iter.ReadArrayCB(func(i *json.Iterator) bool {
		var elem string
		i.ReadVal(&elem)
		row = append(row, elem)
		return true
	}); ok {

		l.Status.RowCount++
		// add header
		if l.Status.RowCount == 1 && !l.Config.NoHeader {
			l.Status.Header = row
		}

		// cell validation
		err = l.lintCell(l.Status.RowCount, row)
		if err != nil {
			iter.Error = err
			return false
		}
		return true
	}
	return true
}

// lintJSON lint JSON file
func (l *LintStruct) lintJSON() error {
	f, err := os.Open(l.Config.File)
	if err != nil {
		return err
	}
	defer f.Close()

	iter := json.Parse(json.ConfigDefault, f, bufio.MaxScanTokenSize)
	switch iter.WhatIsNext() {
	case json.ArrayValue:
		if iter.WhatIsNext() == json.ArrayValue {
			if iter.ReadArrayCB(l.jsonLintRow); iter.Error != nil {
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
