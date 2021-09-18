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

	"d18n/common"

	"github.com/kr/pretty"
)

func TestEmportJSON(t *testing.T) {
	orgCfg := common.Cfg

	common.Cfg.File = common.TestPath + "/test/actor.json"
	common.Cfg.User = ""
	common.Cfg.Limit = 12
	common.Cfg.SkipLines = 1
	common.Cfg.Table = "actor_json"
	common.Cfg.Database = "sakila"
	common.Cfg.Replace = false
	testES.CommonConfig = common.Cfg

	conn, _ := common.NewConnection()
	err := emportJSON(testES, conn)
	if err != nil {
		t.Error(err.Error())
	}

	pretty.Println(testES.Status)

	common.Cfg = orgCfg

}
