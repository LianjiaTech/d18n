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
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/LianjiaTech/d18n/common"
)

// LintStatus returns information about an  record check result
type LintStatus struct {
	Header     []string   // column names
	RowCount   int64      // row count
	CellCount  int        // row cell num
	ErrorCount int64      // error count
	Lint       []LintCode // lint message

	// file extra info
	Size int64 // file size

	TimeCost int64 // total time cost
}

var lintLevels = []string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG"}

type LintStruct struct {
	Config common.Config
	Levels []string
	Rules  map[string]LintCode
	Status LintStatus
}

func NewLintStruct(c common.Config) (*LintStruct, error) {
	var l = &LintStruct{
		Config: c,
		Status: LintStatus{},
		Rules:  make(map[string]LintCode),
	}

	for i, level := range lintLevels {
		if strings.ToUpper(c.LintLevel) == level {
			l.Levels = lintLevels[:i]
			break
		}
	}

	l.initLintRules()

	return l, nil
}

// Lint lint file one time for one file only
func (l *LintStruct) Lint() error {
	lintStartTime := time.Now().UnixNano()

	// run all lint rules about file
	err := l.lintFile(0, []string{})
	if err != nil {
		return err
	}

	// run all lint rules about lines && cells
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(l.Config.File), "."))
	switch suffix {
	case "csv":
		l.Config.Comma = ','
		err = l.lintCSV()
	case "psv":
		l.Config.Comma = '|'
		err = l.lintCSV()
	case "tsv":
		delete(l.Rules, "Whitespace") // tsv not check white space after closed quote
		l.Config.Comma = '\t'
		err = l.lintCSV()
	case "txt":
		delete(l.Rules, "Whitespace") // txt not check white space after closed quote
		l.Config.Comma = ' '
		err = l.lintCSV()
	case "xlsx":
		err = l.lintXlsx()
	case "sql":
		err = l.lintSQL()
	case "json":
		err = l.lintJSON()
	case "html":
		err = l.lintHTML()
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}
	if err != nil {
		return err
	}

	lintEndTime := time.Now().UnixNano()
	l.Status.TimeCost = lintEndTime - lintStartTime
	return err
}

// lintFile ...
// line int64: always 0, no need
// raw []string: always empty, no need
func (l *LintStruct) lintFile(line int64, raw []string) error {

	// test file exist
	info, err := os.Stat(l.Config.File)
	if err != nil {
		return err
	}
	l.Status.Size = info.Size()

	// zero, empty
	if l.Status.Size == 0 {
		return fmt.Errorf("%s file size: 0", l.Config.File)
	}

	// check file size and suffix
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(l.Config.File), "."))
	switch suffix {
	case "xlsx":
		if l.Status.Size > int64(l.Config.ExcelMaxFileSize) {
			err = fmt.Errorf("%s file size: %d, too large", l.Config.File, l.Status.Size)
		}
	case "csv", "txt", "tsv", "psv", "html", "json", "sql":
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}

	return err
}

// lintLine ...
func (l *LintStruct) lintLine(line int64, raw []string) error {
	for _, r := range l.Rules {
		switch r.LintLevel {
		case "line":
			column, err := r.Func(line, raw)
			if err {
				l.Status.ErrorCount++
				rule := r
				rule.Line = line
				rule.Column = column
				l.Status.Lint = append(l.Status.Lint, rule)
				if checkLevelBreak(r) {
					return fmt.Errorf(r.Message)
				}
			}
		}
	}
	return nil
}

// lintCell ...
func (l *LintStruct) lintCell(line int64, raw []string) error {
	for _, r := range l.Rules {
		switch r.LintLevel {
		case "cell":
			column, err := r.Func(line, raw)
			if err {
				l.Status.ErrorCount++
				rule := r
				rule.Line = line
				rule.Column = column
				l.Status.Lint = append(l.Status.Lint, rule)
				if checkLevelBreak(r) {
					return fmt.Errorf(r.Message)
				}
			}
		}
	}
	return nil
}

// ShowStatus check lint status
func (l *LintStruct) ShowStatus() error {
	var err error

	if len(l.Status.Lint) == 0 {
		fmt.Println("ok")
	}

	// format lint status
	for _, s := range l.Status.Lint {
		fmt.Printf("Line: %d, Column: %d, %s: %s\n", s.Line, s.Column, s.Level, s.Message)
	}

	// verbose mode print
	if len(l.Config.Verbose) == 0 {
		return err
	}
	println("")
	println("File Size:", l.Status.Size)

	println("Row Count(Include Header):", l.Status.RowCount,
		"Cell Count:", l.Status.CellCount,
		"Error Count:", l.Status.ErrorCount,
		"Time Cost:", fmt.Sprint(time.Duration(l.Status.TimeCost)*time.Nanosecond))
	return err
}

// lintCellRaggedRows ...
func (l *LintStruct) lintCellRaggedRows(line int64, raw []string) (column int, wrong bool) {
	// lintstatus total column num
	if l.Status.CellCount == 0 && line == 1 {
		l.Status.CellCount = len(raw)
	}
	if len(raw) != l.Status.CellCount {
		return 0, true
	}
	return 0, wrong
}

// lintCellCheckOptions cells less or equal than 1, this may be wrong comma option
func (l *LintStruct) lintCellCheckOptions(line int64, raw []string) (column int, wrong bool) {
	if len(raw) <= 1 {
		return 0, true
	}
	return 0, wrong
}

// lintCellUndeclaredHeader ...
func (l *LintStruct) lintCellUndeclaredHeader(line int64, raw []string) (column int, wrong bool) {
	// fix for sql file
	if line == 1 && !l.Config.NoHeader {
		//raw = l.Status.Header
		if len(l.Status.Header) == 0 {
			return 0, true
		}
		for k, c := range l.Status.Header {
			// empty column name
			if len(c) == 0 {
				return k + 1, true
			}

			// check column name length
			switch l.Config.Server {
			case "mysql":
				if len(c) > 64 {
					return k + 1, true
				}
			case "postgres":
				if len(c) > 63 {
					return k + 1, true
				}
			case "oracle", "sqlserver":
				if len(c) > 128 {
					return k + 1, true
				}
			case "sqlite", "sqlite3":
				if len(c) > 30 {
					return k + 1, true
				}
			}

			// only a-zA-Z, under score, number, not start with number
			r := regexp.MustCompile(`^[a-zA-Z\_][a-zA-Z0-9\_]*`)
			if r.ReplaceAllString(c, "") != "" {
				return k + 1, true
			}
		}
	}
	return 0, wrong
}

// lintCellUnMatchHeader ...
func (l *LintStruct) lintCellUnMatchHeader(line int64, raw []string) (column int, wrong bool) {
	if !l.Config.NoHeader {
		return 0, len(raw) != len(l.Status.Header)
	}
	return 0, wrong
}

// lintCellDupColumnName...
func (l *LintStruct) lintCellDupColumnName(line int64, raw []string) (column int, wrong bool) {
	if line == 1 {
		sort.Strings(raw)
		for k, c := range raw {
			if k > 0 && c == raw[k-1] {
				return k + 1, true
			}
		}
	}
	return 0, wrong
}

// checkLevelBreak ...
func checkLevelBreak(r LintCode) (br bool) {
	for _, level := range lintLevels {
		if level == r.Level {
			br = true
			break
		}
	}
	return br
}
