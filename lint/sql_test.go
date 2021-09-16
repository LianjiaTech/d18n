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

	"d18n/common"

	"github.com/kr/pretty"
)

func TestLintSQLByPingcap(t *testing.T) {
	orgCfg := common.Cfg
	common.Cfg.ANSIQuotes = false
	lintStatus = LintStatus{}
	common.Cfg.File = common.TestPath + "/test/TestSQLLint.right.sql"
	err := lintSQL()
	if err != nil {
		t.Errorf(err.Error())
	}
	pretty.Println(lintStatus)

	lintStatus = LintStatus{}
	common.Cfg.File = common.TestPath + "/test/TestSQLLint.wrong.sql"
	err = lintSQL()
	if err == nil {
		t.Errorf("here should report an error")
	} else {
		pretty.Println(err.Error())
	}
	pretty.Println(lintStatus)

	common.Cfg = orgCfg
}
