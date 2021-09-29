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
	"fmt"
	"testing"

	"github.com/LianjiaTech/d18n/common"
)

var testES *EmportStruct

func init() {
	var err error
	common.InitTestEnv()
	common.TestConfig.Schema = common.TestPath + "/test/schema.txt"
	testES, err = NewEmportStruct(common.TestConfig)
	if err != nil {
		panic(err.Error())
	}
	testES.Status.Header, err = common.TestConfig.ParseSchema()
	fmt.Println(testES.Status.Header)
	if err != nil {
		panic(err.Error())
	}
}

func TestEmportCSV(t *testing.T) {
	orgCfg := common.TestConfig

	common.TestConfig.File = common.TestPath + "/test/actor.csv"
	common.TestConfig.User = ""
	common.TestConfig.Limit = 2
	common.TestConfig.Table = "actor_new"
	common.TestConfig.Replace = true
	common.TestConfig.Comma = ','
	testES.Config = common.TestConfig

	conn, _ := common.TestConfig.NewConnection()
	err := emportCSV(testES, conn)
	if err != nil {
		t.Error(err.Error())
	}

	common.TestConfig = orgCfg

}
