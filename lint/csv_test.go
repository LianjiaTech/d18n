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

	"d18n/common"
)

func init() {
	common.InitTestEnv()
}

func TestLintCSV(t *testing.T) {
	orgCfg := common.Cfg
	common.Cfg.File = common.TestPath + "/test/TestCSVLint.wrong.csv"
	l, _ := NewLintStruct(common.Cfg)
	err := l.lintFile(0, []string{})
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = l.lintCSV()
	if err == nil {
		t.Error("here must have error")
	}

	common.Cfg.File = common.TestPath + "/test/TestCSVLint.right.csv"
	l, _ = NewLintStruct(common.Cfg)
	err = l.lintCSV()
	if err != nil {
		t.Error(err.Error())
	}
	common.Cfg = orgCfg
}

func TestLintCSVUnclosedQuote(t *testing.T) {
	raws := [][]string{
		// right
		{"123,456"},
		{`"abc","def"`},
		{`"a""","def"`},
		{`"a""bc","def"`},
		// wrong
		{`"abc,"abc","112223"`},
		{`"abc"","112223"`},
	}
	for _, raw := range raws[:4] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVUnclosedQuote(1, raw)
		if wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

	for _, raw := range raws[4:] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVUnclosedQuote(1, raw)
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}
}

func TestLintCSVLineBreaks(t *testing.T) {
	orgCfg := common.Cfg
	common.Cfg.LineBreak = "\r\n"
	crlf := `"crlf"` + "\r\n"
	lf := `"lf"` + "\n"

	l, _ := NewLintStruct(common.Cfg)
	column, wrong := l.lintCSVLineBreaks(1, []string{crlf})
	if wrong {
		t.Error(fmt.Sprintf("column: %d", column), crlf)
	}

	l, _ = NewLintStruct(common.Cfg)
	column, wrong = l.lintCSVLineBreaks(1, []string{lf})
	if !wrong {
		t.Error(fmt.Sprintf("column: %d", column), lf)
	}
	common.Cfg = orgCfg
}

func TestLintCSVBlankRows(t *testing.T) {
	raws := [][]string{
		// right
		{"abc"},
		// wrong
		{""},
		{"\t"},
		{"\n"},
		{"\r\n"},
	}
	for _, raw := range raws[:1] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVBlankRows(1, raw)
		if wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

	for _, raw := range raws[1:] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVBlankRows(1, raw)
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}
}

func TestLintCSVCommentRows(t *testing.T) {
	orgCfg := common.Cfg
	// no comment
	l, _ := NewLintStruct(common.Cfg)
	_, wrong := l.lintCSVCommentRows(1, []string{"abc\n"})
	if wrong {
		t.Error("without comment line")
	}

	// comment lines
	common.Cfg.Comments = []string{"//", "#"}
	raws := [][]string{
		{"//abc\n"},
		{"// abc\n"},
		{"// abc\r\n"},
		{"# abc\r\n"},
		{"#abc\r\n"},
	}
	for _, raw := range raws {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVCommentRows(1, []string{string([]byte("//abc\n"))})
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}
	common.Cfg = orgCfg
}

func TestLintCSVCellSpace(t *testing.T) {
	raws := [][]string{
		// right
		{""},
		{"abc"},

		// wrong
		{"abc\n"},
		{"abc\t"},
		{"abc "},
		{"  abc\n"},
		{"\tabc\n"},
	}

	for _, raw := range raws[:2] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVCellSpace(1, raw)
		if wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

	for _, raw := range raws[2:] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVCellSpace(1, raw)
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

}

func TestLintCSVLeadingSpace(t *testing.T) {
	raws := [][]string{
		// right
		{""},
		{"abc"},
		{"abc\n"},
		{"abc\t"},

		// wrong
		{"  abc\n"},
		{"\tabc\n"},
	}

	for _, raw := range raws[:4] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVLeadingSpace(1, raw)
		if wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

	for _, raw := range raws[4:] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVLeadingSpace(1, raw)
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

}

func TestLintCSVWhitespace(t *testing.T) {
	raws := [][]string{
		// right
		{""},
		{`""`},
		{`"abc"`},
		{`"abc"` + "\r"},
		{`"abc"` + "\n"},
		{`"abc"` + "\r\n"},
		{`"abc\n"`},
		{`"abc\t"`},
		{`"abc""","bcd"`},
		{`"abc"" ,","bcd"`},
		{`abc `},
		{"abc\r"},
		{"abc\n"},
		{"abc\r\n"},

		// wrong
		{`"abc" `},
		{`"abc"` + "\t"},
		{`"abc"` + "\v"},
		{`"abc", "def"`},
		{`"abc",` + "\t" + `"def"`},
		{`"abc""" ,"bcd"`},
		{`"abc""", "bcd"`},
		{`"abc" ,"bcd"`},
		{`"abc", "bcd"`},
	}

	for _, raw := range raws[:14] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVWhitespace(1, raw)
		if wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

	for _, raw := range raws[14:] {
		l, _ := NewLintStruct(common.Cfg)
		column, wrong := l.lintCSVWhitespace(1, raw)
		if !wrong {
			t.Error(fmt.Sprintf("column: %d", column), raw)
		}
	}

}
