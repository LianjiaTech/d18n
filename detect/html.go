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

package detect

import (
	"bufio"
	"os"

	"d18n/common"

	"golang.org/x/net/html"
)

func detectHTML() error {
	var err error

	fd, err := os.Open(common.Cfg.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	r := bufio.NewReaderSize(fd, common.Cfg.MaxBufferSize)
	token := html.NewTokenizer(r)

	var row []string
	for {

		t := token.Next()
		if t == html.ErrorToken {
			break
		}

		tag, _ := token.TagName()
		switch t {
		case html.StartTagToken:
			switch string(tag) {
			case "th", "td":
				token.Next()
				row = append(row, html.UnescapeString(string(token.Raw())))
			case "tr":
				detectStatus.Lines++
			}
		case html.EndTagToken:
			switch string(tag) {
			case "tr":
				// check column names
				if detectStatus.Lines == 1 {
					if !common.Cfg.NoHeader && common.Cfg.Schema == "" {
						for _, r := range row {
							detectStatus.Header = append(detectStatus.Header, common.HeaderColumn{Name: r})
						}
					}
					checkHeader(detectStatus.Header)

					// truncate row after new line
					row = []string{}

					if !common.Cfg.NoHeader {
						continue
					}
				}

				// check value
				for j, value := range row {
					detectStatus.Columns[detectStatus.Header[j].Name] = append(detectStatus.Columns[detectStatus.Header[j].Name], checkValue(value)...)
				}

				// truncate row after new line
				row = []string{}
			}
		}

		// SkipLines
		if detectStatus.Lines <= common.Cfg.SkipLines {
			continue
		}
		if common.Cfg.Limit > 0 &&
			(detectStatus.Lines-common.Cfg.SkipLines) > common.Cfg.Limit {
			break
		}

	}

	return err
}
