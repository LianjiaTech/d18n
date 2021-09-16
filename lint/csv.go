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

	"d18n/common"
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
func closedQuoteLineBreak(data []byte) (int, error) {
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
					case common.Cfg.Comma:
						// ",
						closed = !closed
					default:
						head++
						for ; head < tail; head++ {
							if !unicode.IsSpace(rune(data[head])) && rune(data[head]) != common.Cfg.Comma {
								return head, fmt.Errorf("column: %d, have quote error", head-1)
							}
							if rune(data[head]) == '\n' || rune(data[head]) == common.Cfg.Comma {
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
func csvReadLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return
	}

	// Read one csv format row from file
	if i, e := closedQuoteLineBreak(data); i > 0 {
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
func csvReadRow(line string) ([][]string, error) {
	r := csv.NewReader(strings.NewReader(line))
	r.Comma = common.Cfg.Comma
	return r.ReadAll()
}

// lintCSV ...
func lintCSV() error {
	var err error
	//var line int64
	f, err := os.Open(common.Cfg.File)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Buffer([]byte{}, common.Cfg.MaxBufferSize)
	s.Split(csvReadLine)

	for s.Scan() {
		//line++
		//lintStatus.RowCount = line
		lintStatus.RowCount++
		// read lines
		raw := s.Text()

		// line validation
		err = lintLine(lintStatus.RowCount, []string{raw})
		if err != nil {
			return err
		}

		// read cells
		rows, err := csvReadRow(raw)
		if err != nil {
			return err
		}

		// cell validation
		for _, row := range rows {
			if lintStatus.RowCount == 1 && !common.Cfg.NoHeader {
				lintStatus.Header = row
			}
			err = lintCell(lintStatus.RowCount, row)
			if err != nil {
				return err
			}
		}
	}

	if s.Err() != nil {
		// bufio.ErrTooLong 1. raw data too large, 2. missing quotes
		return fmt.Errorf("line: %d, %s", lintStatus.RowCount+1, s.Err().Error())
	}

	return err
}

// lintCSVLineBreaks ...
func lintCSVLineBreaks(line int64, raw []string) (column int, wrong bool) {
	lineBreakLen := len(common.Cfg.LineBreak)
	if lineBreakLen == 0 {
		return 0, wrong
	}

	for _, l := range raw {
		rawLen := len(l)
		if rawLen >= lineBreakLen &&
			l[rawLen-lineBreakLen:rawLen] != string(common.Cfg.LineBreak) {
			return rawLen - lineBreakLen, true
		}
	}
	return 0, wrong
}

// lintCSVBlankRows  ...
func lintCSVBlankRows(line int64, raw []string) (column int, wrong bool) {
	for _, l := range raw {
		if strings.TrimSpace(l) == "" {
			return 0, true
		}
	}
	return 0, wrong
}

// lintCSVUnclosedQuote ...
func lintCSVUnclosedQuote(line int64, raw []string) (column int, wrong bool) {
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
func lintCSVWhitespace(line int64, raw []string) (column int, wrong bool) {
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
func lintCSVCommentRows(line int64, raw []string) (column int, wrong bool) {
	for _, commentChars := range common.Cfg.Comments {
		for _, l := range raw {
			commentLen := len(commentChars)
			if len(l) >= commentLen && commentLen > 0 &&
				l[:commentLen] == commentChars {
				return 0, true
			}
		}
	}
	return 0, false
}

// lintCSVLeadingSpace ...
func lintCSVLeadingSpace(line int64, raw []string) (column int, wrong bool) {
	for _, l := range raw {
		if len(l) > 0 && unicode.IsSpace(rune(l[0])) {
			return column, true
		}
	}
	return column, wrong
}

// lintCSVCellSpace ...
func lintCSVCellSpace(line int64, raw []string) (column int, wrong bool) {
	for _, l := range raw {
		if strings.TrimSpace(l) != l {
			return column, true
		}
		column += len(l)
	}
	return 0, wrong
}
