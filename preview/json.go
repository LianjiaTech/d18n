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
	"fmt"
	"os"

	//"encoding/json"
	json "github.com/json-iterator/go"
)

func previewJSON(p *PreviewStruct) error {
	if p.Config.Preview == 0 {
		return nil
	}

	// use iterator to reader
	fd, err := os.Open(p.Config.File)
	if err != nil {
		return err
	}
	defer fd.Close()

	// read line
	var records []interface{}
	iter := json.Parse(json.ConfigDefault, fd, p.Config.MaxBufferSize)
	if iter.WhatIsNext() == json.ArrayValue {
		for count := 0; iter.ReadArray() && count < p.Config.Preview; count++ {
			ret := iter.Read()
			if iter.Error != nil {
				return iter.Error
			} else {
				records = append(records, ret)
			}
		}
	} else {
		return fmt.Errorf("json file is empty or format not valid")
	}

	// print line
	if out, err := json.MarshalIndent(records, "", "  "); err == nil {
		fmt.Println(string(out))
	} else {
		return err
	}
	return nil
}
