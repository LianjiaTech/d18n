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

// lintFunc ...
// input: line number int64, line raw data []string
// output: error column int, wrong bool
type lintFunc func(line int64, raw []string) (column int, wrong bool)

type LintCode struct {
	Name      string   // lint rule name
	LintLevel string   // file, row, cell
	Level     string   // ERROR, WARN, INFO
	Message   string   // error message
	Column    int      // error column
	Line      int64    // error line
	Func      lintFunc // lint rule function
}

var lintRules map[string]LintCode

func init() {
	lintRules = make(map[string]LintCode)

	lintRules = map[string]LintCode{
		"OK": {
			Name:    "OK",
			Message: "OK.",
		},

		// --------------------------------- ERROR -----------------------------------
		"InvalidEncoding": {
			Name:    "InvalidEncoding",
			Level:   "ERROR",
			Message: "There are any odd characters in a file which could cause encoding errors.",
		},
		"LineBreaks": {
			Name:      "LineBreaks",
			LintLevel: "line",
			Level:     "ERROR",
			Message:   "Line breaks are not the same as define.",
			Func:      lintCSVLineBreaks,
		},
		"RaggedRows": {
			Name:      "RaggedRows",
			LintLevel: "cell",
			Level:     "ERROR",
			Message:   "Rows in the file doesn't have the same number of columns.",
			Func:      lintCellRaggedRows,
		},
		"UnMatchHeader": {
			Name:      "UnMatchHeader",
			LintLevel: "cell",
			Level:     "ERROR",
			Message:   "Header number not match value number.",
			Func:      lintCellUnMatchHeader,
		},
		"UndeclaredHeader": {
			Name:      "UndeclaredHeader",
			LintLevel: "cell",
			Level:     "ERROR",
			Message:   "First line in file can't be used as column names.",
			Func:      lintCellUndeclaredHeader,
		},
		"DupColumnName": {
			Name:      "DuplicateColumnName",
			LintLevel: "cell",
			Level:     "ERROR",
			Message:   "Column names aren't unique.",
			Func:      lintCellDupColumnName,
		},
		"UnclosedQuote": {
			Name:      "UnclosedQuote",
			LintLevel: "line",
			Level:     "ERROR",
			Message:   "There are any unclosed quotes in line.",
			Func:      lintCSVUnclosedQuote,
		},
		// --------------------------------- WARN -----------------------------------
		"CheckOptions": {
			Name:      "CheckOptions",
			LintLevel: "cell",
			Level:     "WARN",
			Message:   "Cells less or equal than 1 .",
			Func:      lintCellCheckOptions,
		},
		"LeadingSpace": {
			Name:      "LeadingSpace",
			LintLevel: "line",
			Level:     "WARN",
			Message:   "Line leading with space.",
			Func:      lintCSVLeadingSpace,
		},
		"CellSpace": {
			Name:      "CellSpace",
			LintLevel: "cell",
			Level:     "WARN",
			Message:   "Cell leading or ending with space.",
			Func:      lintCSVCellSpace,
		},
		"BlankRows": {
			Name:      "BlankRows",
			Level:     "WARN",
			LintLevel: "line",
			Message:   "There are any blank rows.",
			Func:      lintCSVBlankRows,
		},
		"Whitespace": {
			Name:      "Whitespace",
			LintLevel: "line",
			Level:     "WARN",
			Message:   "There is any whitespace between commas and double quotes around cells.",
			Func:      lintCSVWhitespace,
		},
		"CommentRows": {
			Name:      "CommentRows",
			LintLevel: "line",
			Level:     "WARN",
			Message:   "There are any comment rows.",
			Func:      lintCSVCommentRows,
		},

		// --------------------------------- INFO -----------------------------------
	}
}
