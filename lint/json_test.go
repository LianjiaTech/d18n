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
	"testing"

	"github.com/LianjiaTech/d18n/common"

	"github.com/kr/pretty"
)

func TestLintJSON(t *testing.T) {
	orgCfg := common.TestConfig
	levels := lintLevels
	lintLevels = []string{"FATAL", "ERROR"}
	common.TestConfig.NoHeader = true

	common.TestConfig.File = common.TestPath + "/test/TestJSONLint.right.json"
	l, _ := NewLintStruct(common.TestConfig)
	err := l.lintJSON()
	if err != nil {
		t.Errorf(err.Error())
	}
	pretty.Println(l.Status)

	common.TestConfig.File = common.TestPath + "/test/TestJSONLint.wrong.json"
	l, _ = NewLintStruct(common.TestConfig)
	err = l.lintJSON()
	if err == nil {
		t.Errorf("file contain error, but not find")
	}
	pretty.Println(l.Status)

	lintLevels = levels
	common.TestConfig = orgCfg
}
