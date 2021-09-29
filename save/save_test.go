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

package save

import (
	"testing"

	"d18n/common"
)

func TestSave(t *testing.T) {
	orgCfg := common.Cfg
	// check all format file
	files := []string{
		"",
		"stdout",
		common.TestPath + "/test/TestSaveRows.csv",
		// common.TestPath + "/test/TestSaveRows.tsv",
		// common.TestPath + "/test/TestSaveRows.txt",
		// common.TestPath + "/test/TestSaveRows.psv",
		// common.TestPath + "/test/TestSaveRows.sql",
		// common.TestPath + "/test/TestSaveRows.json",
		// common.TestPath + "/test/TestSaveRows.xlsx",
	}
	common.Cfg.Table = "TestSaveRows"

	for _, file := range files {
		common.Cfg.File = file

		// new save struct
		s, err := NewSaveStruct(common.Cfg)
		if err != nil {
			t.Error(err.Error())
		}

		if err := s.Save(); err != nil {
			t.Error(err.Error())
		}
	}
	common.Cfg = orgCfg
}

func TestCheckStatus(t *testing.T) {
	orgCfg := common.Cfg
	s := &SaveStruct{
		Status: saveStatus{
			Lines:    100,
			TimeCost: 1000,
		},
	}

	common.Cfg.Verbose = true
	err := s.ShowStatus()
	if err != nil {
		t.Error(err.Error())
	}
	common.Cfg = orgCfg
}
