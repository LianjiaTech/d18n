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

var detectStatus DetectStatus

func init() {
	var err error
	common.InitTestEnv()

	common.Cfg.Schema = common.TestPath + "/test/schema.txt"

	detectStatus.Header, err = common.ParseSchema()
	if err != nil {
		panic(err.Error())
	}

	detectStatus.Columns = make(map[string][]string, len(detectStatus.Header))

}

func TestDetectCSV(t *testing.T) {
	orgCfg := common.Cfg

	common.Cfg.File = common.TestPath + "/test/actor.csv"
	common.Cfg.User = ""
	common.Cfg.Limit = 10
	common.Cfg.Comma = ','
	d, _ := NewDetectStruct(common.Cfg)
	d.Status = detectStatus

	err := d.detectCSV()
	if err != nil {
		t.Error(err.Error())
	}

	common.Cfg = orgCfg

}
