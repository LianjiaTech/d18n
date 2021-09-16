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

package preview

import (
	"fmt"
	"testing"

	"d18n/common"
)

func TestPreview(t *testing.T) {
	orgCfg := common.Cfg
	files := [][]string{
		{
			common.TestPath + "/test/TestSaveRows.csv",
			common.TestPath + "/test/TestSaveRows.tsv",
			common.TestPath + "/test/TestSaveRows.psv",
			common.TestPath + "/test/TestSaveRows.json",
			common.TestPath + "/test/actor.xlsx",
		},
		{
			"not exist file",
			"",
			"stdout",
		},
	}

	// preview files
	common.Cfg.Preview = 10
	for _, file := range files[0] {
		common.Cfg.File = file
		fmt.Println("# Preview: ", common.Cfg.File)
		err := Preview()
		if err != nil {
			t.Error(err.Error())
		}
	}

	for _, file := range files[1] {
		common.Cfg.File = file
		fmt.Println("# Preview: ", common.Cfg.File)
		err := Preview()
		if err == nil {
			t.Error(err.Error())
		}
	}
	common.Cfg = orgCfg
}
