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

func TestEmportRows(t *testing.T) {
	orgCfg := common.TestConfig
	// check all format file
	files := []string{
		common.TestPath + "/test/actor.csv",
		common.TestPath + "/test/actor.tsv",
		common.TestPath + "/test/actor.txt",
		common.TestPath + "/test/actor.psv",
		common.TestPath + "/test/actor.json",
		common.TestPath + "/test/actor.xlsx",
		common.TestPath + "/test/actor.html",
	}

	common.TestConfig.Schema = common.TestPath + "/test/schema.txt"
	common.TestConfig.User = ""
	common.TestConfig.Limit = 2
	common.TestConfig.Replace = true
	common.TestConfig.Table = "actor"
	common.TestConfig.Database = "sakila"

	conn, err := common.TestConfig.NewConnection()
	if err != nil {
		t.Error(err.Error())
	}

	for i, file := range files {
		common.TestConfig.File = file
		switch i {
		case 0:
			common.TestConfig.Comma = ','
		case 1:
			common.TestConfig.Comma = '\t'
		case 2:
			common.TestConfig.Comma = ' '
		case 3:
			common.TestConfig.Comma = '|'
		}
		fmt.Println(common.TestConfig.File)

		e, err := NewEmportStruct(common.TestConfig)
		if err != nil {
			t.Error(err.Error())
		}

		e.Status.Lines = 0
		err = emportRows(e, conn)
		if err != nil {
			t.Error(err.Error())
		}
	}

	common.TestConfig = orgCfg
}

func TestCheckStatus(t *testing.T) {
	orgCfg := common.TestConfig
	common.TestConfig.Verbose = []bool{true}

	e, err := NewEmportStruct(common.TestConfig)
	if err != nil {
		t.Error(err.Error())
	}

	err = e.ShowStatus()
	if err != nil {
		t.Error(err.Error())
	}

	common.TestConfig = orgCfg
}
