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
)

func previewCSV(p *PreviewStruct) error {
	if p.Config.Preview == 0 {
		return nil
	}

	fd, err := os.Open(p.Config.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	var line int
	s := bufio.NewScanner(fd)
	s.Buffer([]byte{}, p.Config.MaxBufferSize)

	for s.Scan() {
		if line >= p.Config.Preview {
			break
		}
		fmt.Println(s.Text())
		line++
	}

	return s.Err()
}
