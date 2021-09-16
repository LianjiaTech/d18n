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

	"d18n/common"
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

var lintStatus LintStatus

var LintLevels = []string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG"}

// Lint ...
func Lint() error {
	lintStartTime := time.Now().UnixNano()

	lintStatus = LintStatus{}
	for i, l := range LintLevels {
		if strings.ToUpper(common.Cfg.LintLevel) == l {
			LintLevels = LintLevels[:i]
			break
		}
	}

	// run all lint rules about file
	err := lintFile(0, []string{})
	if err != nil {
		return err
	}

	// run all lint rules about lines && cells
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(common.Cfg.File), "."))
	switch suffix {
	case "csv":
		common.Cfg.Comma = ','
		err = lintCSV()
	case "psv":
		common.Cfg.Comma = '|'
		err = lintCSV()
	case "tsv":
		delete(LintRules, "Whitespace") // tsv not check white space after closed quote
		common.Cfg.Comma = '\t'
		err = lintCSV()
	case "txt":
		delete(LintRules, "Whitespace") // txt not check white space after closed quote
		common.Cfg.Comma = ' '
		err = lintCSV()
	case "xlsx":
		err = lintXlsx()
	case "sql":
		err = lintSQL()
	case "json":
		err = lintJSON()
	case "html":
		err = lintHTML()
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}
	if err != nil {
		return err
	}

	lintEndTime := time.Now().UnixNano()
	lintStatus.TimeCost = lintEndTime - lintStartTime
	return err
}

// lintFile ...
// line int64: always 0, no need
// raw []string: always empty, no need
func lintFile(line int64, raw []string) error {

	// test file exist
	info, err := os.Stat(common.Cfg.File)
	if err != nil {
		return err
	}
	lintStatus.Size = info.Size()

	// zero, empty
	if lintStatus.Size == 0 {
		return fmt.Errorf("%s file size: 0", common.Cfg.File)
	}

	// check file size and suffix
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(common.Cfg.File), "."))
	switch suffix {
	case "xlsx":
		if lintStatus.Size > int64(common.Cfg.ExcelMaxFileSize) {
			err = fmt.Errorf("%s file size: %d, too large", common.Cfg.File, lintStatus.Size)
		}
	case "csv", "txt", "tsv", "psv", "html", "json", "sql":
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}

	return err
}

// lintLine ...
func lintLine(line int64, raw []string) error {
	for _, r := range LintRules {
		switch r.LintLevel {
		case "line":
			column, err := r.Func(line, raw)
			if err {
				lintStatus.ErrorCount++
				l := r
				l.Line = line
				l.Column = column
				lintStatus.Lint = append(lintStatus.Lint, r)
				if checkLevelBreak(r) {
					return fmt.Errorf(r.Message)
				}
			}
		}
	}
	return nil
}

// lintCell ...
func lintCell(line int64, raw []string) error {
	for _, r := range LintRules {
		switch r.LintLevel {
		case "cell":
			column, err := r.Func(line, raw)
			if err {
				lintStatus.ErrorCount++
				l := r
				l.Line = line
				l.Column = column
				lintStatus.Lint = append(lintStatus.Lint, l)
				if checkLevelBreak(r) {
					return fmt.Errorf(r.Message)
				}
			}
		}
	}
	return nil
}

// CheckStatus check lint status
func CheckStatus() error {
	var err error

	if len(lintStatus.Lint) == 0 {
		fmt.Println("ok")
	}

	// format lint status
	for _, l := range lintStatus.Lint {
		fmt.Printf("Line: %d, Column: %d, %s: %s\n", l.Line, l.Column, l.Level, l.Message)
	}

	// verbose mode print
	if !common.Cfg.Verbose {
		return err
	}
	println("")
	println("File Size:", lintStatus.Size)

	println("Row Count(Include Header):", lintStatus.RowCount,
		"Cell Count:", lintStatus.CellCount,
		"Error Count:", lintStatus.ErrorCount,
		"Time Cost:", fmt.Sprint(time.Duration(lintStatus.TimeCost)*time.Nanosecond))
	return err
}

// lintCellRaggedRows ...
func lintCellRaggedRows(line int64, raw []string) (column int, wrong bool) {
	// lintstatus total column num
	if lintStatus.CellCount == 0 && line == 1 {
		lintStatus.CellCount = len(raw)
	}
	if len(raw) != lintStatus.CellCount {
		return 0, true
	}
	return 0, wrong
}

// lintCellCheckOptions cells less or equal than 1, this may be wrong comma option
func lintCellCheckOptions(line int64, raw []string) (column int, wrong bool) {
	if len(raw) <= 1 {
		return 0, true
	}
	return 0, wrong
}

// lintCellUndeclaredHeader ...
func lintCellUndeclaredHeader(line int64, raw []string) (column int, wrong bool) {
	// fix for sql file
	if line == 1 && !common.Cfg.NoHeader {
		//raw = lintStatus.Header
		if len(lintStatus.Header) == 0 {
			return 0, true
		}
		for k, c := range lintStatus.Header {
			// empty column name
			if len(c) == 0 {
				return k + 1, true
			}

			// check column name length
			switch common.Cfg.Server {
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
			case "sqlite":
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
func lintCellUnMatchHeader(line int64, raw []string) (column int, wrong bool) {
	if !common.Cfg.NoHeader {
		return 0, len(raw) != len(lintStatus.Header)
	}
	return 0, wrong
}

// lintCellDupColumnName...
func lintCellDupColumnName(line int64, raw []string) (column int, wrong bool) {
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
	for _, l := range LintLevels {
		if l == r.Level {
			br = true
			break
		}
	}
	return br
}
