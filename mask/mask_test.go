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
	"fmt"
	"testing"

	"github.com/LianjiaTech/d18n/common"
)

// TestMask test FakeXXXX
func TestMask(t *testing.T) {
	m, err := NewMaskStruct(common.TestConfig.Mask)
	if err != nil {
		t.Error(err.Error())
	}

	m.Config = map[string]maskRule{
		"col1": {
			MaskFunc: "fake",
			Args:     []string{"name"},
		},
		"col2": {
			MaskFunc: "smokeleft",
			Args:     []string{"3", "*"},
		},
		"col3": {
			MaskFunc: "json",
		},
		"phoneno": {
			MaskFunc: "fake",
			Args:     []string{"phone"}, // phone return string, not int, so it doesn't mask value
		},
		"firstname": {
			MaskFunc: "shuffle",
		},
	}

	cases := [][]interface{}{
		{"col", 1},
		{"col1", "hello world"},
		{"col2", "hello earth"},
		{"col3", `{"foo":1,"bar":2,"baz":[3,4],"phoneNo":13888880000, "newField":"test", "userInfo":{"firstname":"Kritchat", "lastname": "Rojanaphruk"}}`},
	}

	for _, c := range cases {
		ret, err := m.Mask(c[0].(string), c[1])
		if err != nil {
			t.Error(m.Config[c[0].(string)], err.Error())
		}
		fmt.Println(c[0], m.Config[c[0].(string)], ret)
	}
}

func ExampleMask() {
	m, _ := NewMaskStruct(common.TestConfig.Mask)
	m.Config = map[string]maskRule{
		"col1": {
			MaskFunc: "crc32",
			Args:     []string{},
		},
	}

	fmt.Println(m.Mask("col", "hello world"))
	fmt.Println(m.Mask("col1", "hello world"))
	// Output:
	// hello world <nil>
	// 0d4a1185 <nil>
}
