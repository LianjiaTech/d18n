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

package lint

import (
	"fmt"
	"testing"

	"github.com/LianjiaTech/d18n/common"
)

func TestLint(t *testing.T) {
	orgCfg := common.TestConfig
	// check all format file
	files := [][]string{
		// right
		{
			common.TestPath + "/test/TestCSVLint.right.csv",
			common.TestPath + "/test/TestCSVLint.right.tsv",
			common.TestPath + "/test/TestCSVLint.right.psv",
			common.TestPath + "/test/TestCSVLint.right.txt",
			common.TestPath + "/test/TestJSONLint.right.json",
			common.TestPath + "/test/TestXLSXLint.right.xlsx",
			common.TestPath + "/test/TestSQLLint.right.sql",
			common.TestPath + "/test/TestHTMLLint.right.html",
		},
		// wrong
		{
			"",
			"stdout",
			common.TestPath + "/test/TestCSVLint.wrong.csv",
			common.TestPath + "/test/TestSaveRows.tsv",
			common.TestPath + "/test/TestSaveRows.txt",
			common.TestPath + "/test/TestSaveRows.psv",
			common.TestPath + "/test/TestJSONLint.wrong.json",
			common.TestPath + "/test/TestXLSXLint.wrong.xlsx",
			common.TestPath + "/test/TestSQLLint.wrong.sql",
			common.TestPath + "/test/TestHTMLLint.wrong.html",
		},
	}

	// right
	for _, file := range files[0] {
		common.TestConfig.File = file
		l, _ := NewLintStruct(common.TestConfig)
		err := l.Lint()
		if err != nil {
			t.Error(file, err.Error())
		}
	}

	// wrong
	for _, file := range files[1] {
		common.TestConfig.File = file
		l, _ := NewLintStruct(common.TestConfig)
		err := l.Lint()
		if err == nil {
			t.Error(file, "should wrong")
		}
	}
	common.TestConfig = orgCfg
}

func TestLintLine(t *testing.T) {
	raws := [][]string{
		// right
		{"abc", "def"},
		// wrong
		{""},     // blank row
		{"\n"},   // blank row
		{"\r\n"}, // blank row
		{`"abc`}, // unclosed quote
	}
	for _, raw := range raws[:1] {
		l, _ := NewLintStruct(common.TestConfig)
		err := l.lintLine(1, raw)
		if err != nil {
			t.Error(err.Error())
		}
	}

	for _, raw := range raws[1:] {
		l, _ := NewLintStruct(common.TestConfig)
		err := l.lintLine(1, raw)
		if err == nil {
			t.Error("should get error")
		}
	}
}

func TestLintCellRaggedRows(t *testing.T) {
	raw := []string{`"abc`, `"abcc"`, `"112223"`}
	l, _ := NewLintStruct(common.TestConfig)
	l.Status.CellCount = 0
	column, wrong := l.lintCellRaggedRows(1, raw)
	if wrong {
		t.Error(fmt.Sprintf("column: %d", column), raw)
	}

	column, wrong = l.lintCellRaggedRows(2, raw)
	if wrong {
		t.Error(fmt.Sprintf("column: %d", column), raw)
	}

	l.Status.CellCount = 2
	column, wrong = l.lintCellRaggedRows(3, raw)
	if !wrong {
		t.Error(fmt.Sprintf("column: %d", column), raw)
	}

	l.Status.CellCount = 0
}

func TestLintCellCheckOptions(t *testing.T) {
	raws := [][]string{
		// right
		{`abc`, `abcc`},
		// wrong
		{`abc`},
		{""},
		{},
	}
	for _, raw := range raws[:1] {
		l, _ := NewLintStruct(common.TestConfig)
		column, wrong := l.lintCellCheckOptions(1, raw)
		if wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

	for _, raw := range raws[1:] {
		l, _ := NewLintStruct(common.TestConfig)
		column, wrong := l.lintCellCheckOptions(1, raw)
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}
}

func TestLintCellUndeclaredHeader(t *testing.T) {
	raws := [][]string{
		// right
		{"abc", "ddd_ddd", "abc123", "Abc123", "_a"},                         // ok
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, // 64 characters

		// wrong
		{"112223"},  // pure number
		{"ddd-ddd"}, // -
		{"112&223"}, // &
		{"112@223"}, // @
		{"a.b"},     // .
		{""},        // empty
		{"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"}, // 65 characters
	}

	// right
	for _, raw := range raws[:2] {
		l, _ := NewLintStruct(common.TestConfig)
		l.Status.Header = raws[0]
		column, wrong := l.lintCellUndeclaredHeader(1, raw)
		if wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}
	// wrong
	for _, raw := range raws[2:] {
		l, _ := NewLintStruct(common.TestConfig)
		l.Status.Header = raws[2]
		l.Status.Header = []string{"112223"}
		column, wrong := l.lintCellUndeclaredHeader(1, raw)
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}
}

func TestLintCellDupColumnName(t *testing.T) {
	raws := [][]string{
		// right
		{"a", "b"}, // ok

		// wrong
		{"a", "a"},      // duplicate
		{"a", "b", "a"}, // duplicate
	}

	// right
	for _, raw := range raws[:1] {
		l, _ := NewLintStruct(common.TestConfig)
		column, wrong := l.lintCellDupColumnName(1, raw)
		if wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

	// wrong
	for _, raw := range raws[1:] {
		l, _ := NewLintStruct(common.TestConfig)
		column, wrong := l.lintCellDupColumnName(1, raw)
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}
}
