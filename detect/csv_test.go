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
	"testing"

	"d18n/common"
)

// test detect status for Header， Columns init
var detectTestStatus DetectStatus

func init() {
	var err error
	common.InitTestEnv()

	common.TestConfig.Schema = common.TestPath + "/test/schema.txt"

	detectTestStatus.Header, err = common.TestConfig.ParseSchema()
	if err != nil {
		panic(err.Error())
	}
	detectTestStatus.Columns = make(map[string][]string, len(detectTestStatus.Header))
}

func TestDetectCSV(t *testing.T) {
	orgCfg := common.TestConfig

	common.TestConfig.File = common.TestPath + "/test/actor.csv"
	common.TestConfig.User = ""
	common.TestConfig.Limit = 10
	common.TestConfig.Comma = ','
	d, err := NewDetectStruct(common.TestConfig)
	if err != nil {
		t.Errorf(err.Error())
	}
	d.Status = detectTestStatus

	err = d.detectCSV()
	if err != nil {
		t.Error(err.Error())
	}

	common.TestConfig = orgCfg
}
