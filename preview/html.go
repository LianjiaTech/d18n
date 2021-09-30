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

package preview

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func previewHTML(p *PreviewStruct) error {
	if p.Config.Preview == 0 {
		return nil
	}

	file, err := os.Open(p.Config.File)
	if err != nil {
		return err
	}
	defer file.Close()

	r := bufio.NewReaderSize(file, p.Config.MaxBufferSize)
	token := html.NewTokenizer(r)

	var line int
	for {
		if line >= p.Config.Preview {
			break
		}

		t := token.Next()
		if t == html.ErrorToken {
			break
		}

		tag, _ := token.TagName()
		switch t {
		case html.StartTagToken:
			switch string(tag) {
			case "td", "th":
				token.Next()
				fmt.Print(html.UnescapeString(string(token.Raw())), "\t")
			}
		case html.EndTagToken:
			switch string(tag) {
			case "tr":
				fmt.Printf("\n")
				line++
			}
		}
	}

	return err
}
