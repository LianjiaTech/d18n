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
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// RFC4180
// Common Format and MIME Type for Comma-Separated Values (CSV) Files
// https://datatracker.ietf.	org/doc/html/rfc4180

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// closedQuoteLineBreak quote close and line break
func (l *LintStruct) closedQuoteLineBreak(data []byte) (int, error) {
	var closed = true
	var head int
	var tail = len(data)
	for ; head < tail; head++ {
		current := data[head]
		if current == '"' {
			if closed {
				closed = !closed
			} else {
				if head+1 < tail {
					next := data[head+1]
					switch rune(next) {
					case '"':
						// ""
						head++
					case l.CommonConfig.Comma:
						// ",
						closed = !closed
					default:
						head++
						for ; head < tail; head++ {
							if !unicode.IsSpace(rune(data[head])) && rune(data[head]) != l.CommonConfig.Comma {
								return head, fmt.Errorf("column: %d, have quote error", head-1)
							}
							if rune(data[head]) == '\n' || rune(data[head]) == l.CommonConfig.Comma {
								closed = !closed
								break
							}
						}
						head--
					}
				}
			}
		}
		if closed && current == '\n' {
			return head, nil
		}
	}
	return -1, nil
}

// csvReadLine bufio.Scan() SplitFunc
func (l *LintStruct) csvReadLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return
	}

	// Read one csv format row from file
	if i, e := l.closedQuoteLineBreak(data); i > 0 {
		advance = i + 1
		token = data[0 : i+1]
		err = e
		return
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		advance = len(data)
		token = dropCR(data)
		return
	}

	// Request more data.
	return
}

// csvReadRow convert raw line into cells list, reuse csv Reader
func (l *LintStruct) csvReadRow(line string) ([][]string, error) {
	r := csv.NewReader(strings.NewReader(line))
	r.Comma = l.CommonConfig.Comma
	return r.ReadAll()
}

// lintCSV ...
func (l *LintStruct) lintCSV() error {
	var err error
	//var line int64
	f, err := os.Open(l.CommonConfig.File)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer([]byte{}, l.CommonConfig.MaxBufferSize)
	s.Split(l.csvReadLine)

	for s.Scan() {
		//line++
		//l.Status.RowCount = line
		l.Status.RowCount++
		// read lines
		raw := s.Text()

		// line validation
		err = l.lintLine(l.Status.RowCount, []string{raw})
		if err != nil {
			return err
		}

		// read cells
		rows, err := l.csvReadRow(raw)
		if err != nil {
			return err
		}

		// cell validation
		for _, row := range rows {
			if l.Status.RowCount == 1 && !l.CommonConfig.NoHeader {
				l.Status.Header = row
			}
			err = l.lintCell(l.Status.RowCount, row)
			if err != nil {
				return err
			}
		}
	}

	if s.Err() != nil {
		// bufio.ErrTooLong 1. raw data too large, 2. missing quotes
		return fmt.Errorf("line: %d, %s", l.Status.RowCount+1, s.Err().Error())
	}

	return err
}

// lintCSVLineBreaks ...
func (l *LintStruct) lintCSVLineBreaks(line int64, raw []string) (column int, wrong bool) {
	lineBreakLen := len(l.CommonConfig.LineBreak)
	if lineBreakLen == 0 {
		return 0, wrong
	}

	for _, buf := range raw {
		rawLen := len(buf)
		if rawLen >= lineBreakLen &&
			buf[rawLen-lineBreakLen:rawLen] != string(l.CommonConfig.LineBreak) {
			return rawLen - lineBreakLen, true
		}
	}
	return 0, wrong
}

// lintCSVBlankRows  ...
func (l *LintStruct) lintCSVBlankRows(line int64, raw []string) (column int, wrong bool) {
	for _, buf := range raw {
		if strings.TrimSpace(buf) == "" {
			return 0, true
		}
	}
	return 0, wrong
}

// lintCSVUnclosedQuote ...
func (l *LintStruct) lintCSVUnclosedQuote(line int64, raw []string) (column int, wrong bool) {
	var closed = true
	for _, chars := range raw {
		for _, c := range chars {
			if c == '"' {
				closed = !closed
			}
		}
		column += len(chars)
	}
	wrong = !closed

	return column, wrong
}

// lintCSVWhitespace check whitespace after closed quote
func (l *LintStruct) lintCSVWhitespace(line int64, raw []string) (column int, wrong bool) {
	var closed = true
	var hasQuote bool
	for _, chars := range raw {
		for _, c := range chars {
			if c == '"' {
				hasQuote = true
				closed = !closed
			}
			if closed && hasQuote && unicode.IsSpace(c) {
				switch c {
				case '\r', '\n':
					continue
				}
				return column, true
			}
		}
		column += len(chars)
	}
	return column, wrong
}

// lintCommentRows ...
func (l *LintStruct) lintCSVCommentRows(line int64, raw []string) (column int, wrong bool) {
	for _, commentChars := range l.CommonConfig.Comments {
		for _, buf := range raw {
			commentLen := len(commentChars)
			if len(buf) >= commentLen && commentLen > 0 &&
				buf[:commentLen] == commentChars {
				return 0, true
			}
		}
	}
	return 0, false
}

// lintCSVLeadingSpace ...
func (l *LintStruct) lintCSVLeadingSpace(line int64, raw []string) (column int, wrong bool) {
	for _, buf := range raw {
		if len(buf) > 0 && unicode.IsSpace(rune(buf[0])) {
			return column, true
		}
	}
	return column, wrong
}

// lintCSVCellSpace ...
func (l *LintStruct) lintCSVCellSpace(line int64, raw []string) (column int, wrong bool) {
	for _, buf := range raw {
		if strings.TrimSpace(buf) != buf {
			return column, true
		}
		column += len(buf)
	}
	return 0, wrong
}
