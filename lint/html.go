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
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func (l *LintStruct) lintHTML() error {
	var err error
	fd, err := os.Open(l.Config.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	r := bufio.NewReaderSize(fd, l.Config.MaxBufferSize)
	token := html.NewTokenizer(r)
	// max depth is 2 in table tag
	var depth int
	var row []string
	for {
		t := token.Next()
		if t == html.ErrorToken {
			break
		}
		tag, _ := token.TagName()
		if token.Err() != nil {
			return err
		}
		// only lint table tag
		if string(tag) == "table" && t == html.StartTagToken {
			for !((string(tag) == "table" && t == html.EndTagToken) || t == html.ErrorToken) {
				t = token.Next()
				tag, _ = token.TagName()
				if t == html.CommentToken {
					continue
				}

				if token.Err() != nil {
					return err
				}
				// html parser depth record
				// <tr> 、<th> 、<td>  depth +1
				if t == html.StartTagToken && (string(tag) == "tr" || string(tag) == "th" || string(tag) == "td") {
					if string(tag) == "th" {
						l.Config.NoHeader = false
					}
					depth++
				}
				// </tr> 、</th> 、</td> depth -1
				if t == html.EndTagToken && (string(tag) == "tr" || string(tag) == "th" || string(tag) == "td") {
					depth--
				}

				// 1、 invalid tag lint,table tag only support tr, td, th table comment tag
				// 2、tag end lint, e.g., ...<tr><th>a</tr>...
				// 3、tag start lint, e.g., ...<tr>a</th></tr>...
				if err = l.lintTag(depth, tag, t); err != nil {
					return err
				}

				switch string(tag) {
				case "th", "td":
					if depth == 2 {
						t = token.Next()
						row = append(row, string(token.Raw()))
					}
				}
				// illegal character between two tag ,e.g., ...</tr> abc</tr>...
				if depth == 0 && t == html.TextToken && len(strings.TrimSpace(string(token.Raw()))) > 0 {
					return fmt.Errorf("line %d illegal character between two tag %s", l.Status.RowCount, string(token.Raw()))
				}
				if depth == 0 && string(tag) != "table" && t != html.TextToken {
					l.Status.RowCount++
					if l.Status.RowCount == 1 && !l.Config.NoHeader {
						l.Status.Header = row
					}
					err = l.lintCell(l.Status.RowCount, row)
					if err != nil {
						return err
					}
					row = []string{}
				}
			}
		}
	}
	return err
}

func (l *LintStruct) lintTag(depth int, tag []byte, t html.TokenType) error {
	// invalid tag lint,table tag only support tr, td, th table comment tag
	if !(string(tag) == "th" || string(tag) == "tr" || string(tag) == "td" || string(tag) == "table" || t == html.TextToken) {
		return fmt.Errorf("line %d exist invalid tag: %s", l.Status.RowCount, string(tag))
	}
	// tag end lint, e.g., ...<tr><th>a</tr>...
	if depth > 2 {
		return fmt.Errorf("line %d miss end tag: </%s>", l.Status.RowCount, string(tag))
	}
	// tag start lint, e.g., ...<tr>a</th></tr>...
	if depth < 0 {
		return fmt.Errorf("line %d miss start tag: </%s>", l.Status.RowCount, string(tag))
	}
	return nil
}
