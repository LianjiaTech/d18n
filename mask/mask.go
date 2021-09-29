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

package mask

import (
	"embed"
	"fmt"
	"strings"

	"github.com/LianjiaTech/d18n/common"
)

//go:embed corpus
var corpusFS embed.FS

type MaskStruct struct {
	Config maskConfig // mask config
}

func NewMaskStruct(file string) (*MaskStruct, error) {
	var m *MaskStruct
	mc, err := ParseMaskConfig(file)
	if err != nil {
		return m, err
	}

	m = &MaskStruct{
		Config: mc,
	}
	return m, nil
}

// Mask mask column data for SQL query
// name: column name, case insensitive
// value: column value
func (m *MaskStruct) Mask(name string, value interface{}) (ret string, err error) {
	// column name case insensitive
	name = strings.ToLower(name)

	// check mask config
	if _, ok := m.Config[name]; !ok {
		return fmt.Sprint(value), nil
	}
	if _, ok := maskFuncs[m.Config[name].MaskFunc]; !ok {
		return fmt.Sprint(value), fmt.Errorf(common.WrongMaskFunc)
	}

	// concat mask args
	var args []interface{}
	// generate fake data no need origin value
	if !strings.HasPrefix(m.Config[name].MaskFunc, "fake") {
		args = append(args, value)
	}
	for _, arg := range m.Config[name].Args {
		args = append(args, arg)
	}

	// run mask function
	mask := maskFuncs[m.Config[name].MaskFunc]
	return mask(args...)
}

// MaskRow mask data for read file
func (m *MaskStruct) MaskRow(header []common.HeaderColumn, row []string) (ret []string, err error) {
	if len(header) != len(row) {
		return ret, fmt.Errorf(common.WrongColumnsCnt)
	}

	ret = make([]string, len(header))
	for i, h := range header {
		ret[i], err = m.Mask(h.Name, row[i])
		if err != nil {
			return
		}
	}
	return
}
