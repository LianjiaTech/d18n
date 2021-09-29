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
	"database/sql"
	"testing"

	"d18n/common"
)

func init() {
	common.InitTestEnv()
}

func TestSaveRows2ASCII(t *testing.T) {
	// new save struct
	s, err := NewSaveStruct(common.TestConfig)
	if err != nil {
		t.Error(err.Error())
	}

	rows, err := testQueryRows(s)
	if err != nil {
		t.Error(err.Error())
	}
	saveRows2ASCII(s, rows)
}

func testQueryRows(s *SaveStruct) (*sql.Rows, error) {
	// QueryRows
	rows, err := s.Config.QueryRows()
	if err != nil {
		return rows, err
	}

	// get column header
	s.Status.Header, err = rows.ColumnTypes()
	if err != nil {
		return rows, err
	}
	return rows, err
}
