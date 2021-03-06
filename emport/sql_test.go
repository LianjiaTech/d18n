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

package emport

import (
	"testing"

	"github.com/LianjiaTech/d18n/common"

	"github.com/kr/pretty"
)

func TestEmportSQL(t *testing.T) {
	orgCfg := common.TestConfig

	common.TestConfig.File = common.TestPath + "/test/actor.sql"
	common.TestConfig.User = ""
	common.TestConfig.Limit = 12
	common.TestConfig.SkipLines = 1
	common.TestConfig.Table = "actor_sql"
	common.TestConfig.Database = "sakila"
	testES.Config = common.TestConfig
	testES.Status.Lines = 0

	conn, _ := common.TestConfig.NewConnection()
	err := emportSQL(testES, conn)
	if err != nil {
		t.Error(err.Error())
	}
	pretty.Println(testES.Status)

	common.TestConfig = orgCfg
}
