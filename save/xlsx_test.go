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

	"github.com/LianjiaTech/d18n/common"
)

func TestSaveRows2ELSX(t *testing.T) {
	orgCfg := common.TestConfig
	common.TestConfig.File = common.TestPath + "/test/TestSaveRows2ELSX.xlsx"

	// new save struct
	s, err := NewSaveStruct(common.TestConfig)
	if err != nil {
		t.Error(err.Error())
	}

	if err := s.Save(); err != nil {
		t.Error(err.Error())
	}

	common.TestConfig = orgCfg
}
