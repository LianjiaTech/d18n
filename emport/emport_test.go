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

	"d18n/common"
)

func TestEmportRows(t *testing.T) {
	orgCfg := common.Cfg
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

	common.Cfg.Schema = common.TestPath + "/test/schema.txt"
	common.Cfg.User = ""
	common.Cfg.Limit = 2
	common.Cfg.Replace = true
	common.Cfg.Table = "actor"
	common.Cfg.Database = "sakila"

	conn, err := common.NewConnection()
	if err != nil {
		t.Error(err.Error())
	}

	for i, file := range files {
		common.Cfg.File = file
		switch i {
		case 0:
			common.Cfg.Comma = ','
		case 1:
			common.Cfg.Comma = '\t'
		case 2:
			common.Cfg.Comma = ' '
		case 3:
			common.Cfg.Comma = '|'
		}
		fmt.Println(common.Cfg.File)

		e, err := NewEmportStruct(common.Cfg)
		if err != nil {
			t.Error(err.Error())
		}

		e.Status.Lines = 0
		err = emportRows(e, conn)
		if err != nil {
			t.Error(err.Error())
		}
	}

	common.Cfg = orgCfg
}

func TestCheckStatus(t *testing.T) {
	orgCfg := common.Cfg
	common.Cfg.Verbose = true

	e, err := NewEmportStruct(common.Cfg)
	if err != nil {
		t.Error(err.Error())
	}

	err = e.ShowStatus()
	if err != nil {
		t.Error(err.Error())
	}

	common.Cfg = orgCfg
}
