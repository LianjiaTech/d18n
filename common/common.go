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

package common

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
)

var TestConfig Config
var TestPath string

func InitTestEnv() {
	_, filename, _, _ := runtime.Caller(0)
	TestPath = filepath.Dir(filepath.Dir(filename))

	Cfg = Config{
		Server:   "mysql",
		User:     "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     "3306",
		Database: "",
		Charset:  "utf8mb4",

		Query: `select
		'test', '中文', # test string unicode
		1, 0.4, # test numeric
		0.0000000000000000000000000000000000000000000000000000000000000000000000001, # 0, decision
		1e+2, 1E+2, # scientific notation
		'NULL', NULL, # test NULL value
		'double quote"', 'single quote\'' # test quotes
		`,

		NULLString:     "NULL",
		MaxBufferSize:  bufio.MaxScanTokenSize,
		Comma:          ',',
		Comments:       []string{"#", "--"},
		SkipLines:      1,
		ExtendedInsert: 1,

		ExcelMaxFileSize: DefaultExcelMaxFileSize,

		Cipher: TestPath + "/test/cipher.yaml",
	}
}

// PanicIfError panic if error
func PanicIfError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// StringUnique string list 去重，同时也会去除空格
func StringUnique(stringSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			if entry == "" {
				continue
			}
			list = append(list, entry)
		}
	}
	return list
}
