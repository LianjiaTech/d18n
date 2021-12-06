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
	"fmt"
	"os"
	"testing"

	"github.com/kr/pretty"
)

func TestParseDefaultsExtraFile(t *testing.T) {
	orgCfg := TestConfig
	err := parseDefaultsExtraFile(TestPath+"/test/my.cnf", &TestConfig)
	if err != nil {
		t.Error(err.Error())
	}
	pretty.Println(TestConfig.User, TestConfig.Password, TestConfig.Charset)

	TestConfig = orgCfg
}

func TestParseLoginPath(t *testing.T) {
	orgCfg := TestConfig

	orgLoginFile := os.Getenv("MYSQL_TEST_LOGIN_FILE")

	err := os.Setenv("MYSQL_TEST_LOGIN_FILE", TestPath+"/test/.mylogin.cnf")
	if err != nil {
		t.Error(err.Error())
	}

	err = parseLoginPath("d18n", &TestConfig)
	if err != nil {
		t.Error(err.Error())
	}
	pretty.Println(TestConfig.User, TestConfig.Password, TestConfig.Charset)

	os.Setenv("MYSQL_TEST_LOGIN_FILE", orgLoginFile)
	TestConfig = orgCfg
}

func Example_parseCommaFlag() {
	fmt.Println(parseCommaFlag(""))
	fmt.Println(parseCommaFlag(" "))
	fmt.Println(parseCommaFlag("col1, col2"))
	fmt.Println(parseCommaFlag("col1,col2"))
	// Output:
	// []
	// []
	// [col1 col2]
	// [col1 col2]
}

func ExampleParseSchema() {
	orgCfg := TestConfig

	TestConfig.Database = "sakila"
	TestConfig.Table = "actor"
	fmt.Println(TestConfig.ParseSchema())
	// Output:
	// [{actor_id uint16 SMALLINT} {first_name RawBytes VARCHAR} {last_name RawBytes VARCHAR} {last_update NullTime TIMESTAMP}] <nil>

	TestConfig = orgCfg
}
