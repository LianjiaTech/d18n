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
	"testing"

	"github.com/kr/pretty"
)

func TestParseDefaultsExtraFile(t *testing.T) {
	orgCfg := Cfg
	err := parseDefaultsExtraFile(TestPath + "/test/my.cnf")
	if err != nil {
		t.Error(err.Error())
	}
	pretty.Println(Cfg.User, Cfg.Password, Cfg.Charset)
	Cfg = orgCfg
}

func TestParseMaskConfig(t *testing.T) {
	orgCfg := Cfg
	orgMask := MaskConfig

	Cfg.Mask = TestPath + "/test/mask.csv"
	err := ParseMaskConfig()
	if err != nil {
		t.Error(err.Error())
	}
	pretty.Println(MaskConfig)

	MaskConfig = orgMask
	Cfg = orgCfg
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
	orgCfg := Cfg

	Cfg.Database = "sakila"
	Cfg.Table = "actor"
	fmt.Println(ParseSchema())
	// Output:
	// [{actor_id uint16 SMALLINT} {first_name RawBytes VARCHAR} {last_name RawBytes VARCHAR} {last_update NullTime TIMESTAMP}] <nil>

	Cfg = orgCfg
}
