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
	"testing"

	"github.com/kr/pretty"
)

func ExampleParseSelectFields() {
	orgCfg := TestConfig
	TestConfig.Parser = "mysql"
	TestConfig.Query = `select "a", "1", 1, c1 as c, c2, (c2 ), t. c3, count(c4) as cnt, sum(c5+c6) as s, * from db.tb as t where c7 = 2 union select * from t2`
	pretty.Println(TestConfig.ParseSelectFields())
	TestConfig.Parser = "postgres"
	TestConfig.Query = `select "a", "1", 1, c1 as c, c2, (c2 ), t. c3, count(c4) as cnt, sum(c5+c6) as s, * from db.tb as t where c7 = 2 union select * from t2`
	pretty.Println(TestConfig.ParseSelectFields())
	// Output:
	// common.SelectFields{
	//     Fields: {
	//         {Database:"", Table:"", Name:"a", As:""},
	//         {Database:"", Table:"", Name:"1", As:""},
	//         {Database:"", Table:"", Name:"1", As:""},
	//         {Database:"", Table:"", Name:"c1", As:"c"},
	//         {Database:"", Table:"", Name:"c2", As:""},
	//         {Database:"", Table:"", Name:"c2", As:""},
	//         {Database:"", Table:"t", Name:"c3", As:""},
	//         {Database:"", Table:"", Name:"*", As:""},
	//         {Database:"", Table:"", Name:"*", As:""},
	//     },
	// } nil
	// common.SelectFields{
	//     Fields: {
	//         {Database:"", Table:"", Name:"a", As:""},
	//         {Database:"", Table:"", Name:"1", As:""},
	//         {Database:"", Table:"", Name:"1", As:""},
	//         {Database:"", Table:"", Name:"c1", As:"c"},
	//         {Database:"", Table:"", Name:"c2", As:""},
	//         {Database:"", Table:"", Name:"c2", As:""},
	//         {Database:"", Table:"t", Name:"c3", As:""},
	//         {Database:"", Table:"", Name:"*", As:""},
	//         {Database:"", Table:"", Name:"*", As:""},
	//     },
	// } nil

	TestConfig = orgCfg
}

func ExampleParseSelectTables() {
	orgCfg := TestConfig
	TestConfig.Parser = "mysql"
	TestConfig.Query = `select "1", 1, c1 as c, c2, t.c3, * from db.tb as t`
	pretty.Println(TestConfig.ParseSelectTables())
	TestConfig.Parser = "postgres"
	TestConfig.Query = `select "1", 1, c1 as c, c2, t.c3, * from db.tb as t`
	pretty.Println(TestConfig.ParseSelectTables())
	// Output:
	// common.SelectTables{
	//     Tables: {
	//         {Database:"db", Table:"tb"},
	//     },
	// } nil
	// common.SelectTables{
	//     Tables: {
	//         {Database:"db", Table:"tb"},
	//     },
	// } nil
	TestConfig = orgCfg
}

func ExampleParseSelectFuncs() {
	orgCfg := TestConfig
	TestConfig.Parser = "mysql"
	TestConfig.Query = `select binary "a", upper("abc"), count(*), sum(c1) from tb`
	pretty.Println(TestConfig.ParseSelectFuncs())
	TestConfig.Parser = "postgres"
	TestConfig.Query = `select upper('abc'), count(*), sum(c1) from tb`
	pretty.Println(TestConfig.ParseSelectFuncs())
	// Output:
	// common.SelectFuncs{
	//     Funcs: {"binary", "upper", "count", "sum"},
	// } nil
	// common.SelectFuncs{
	//     Funcs: {"upper", "count", "sum"},
	// } nil
	TestConfig = orgCfg
}

func TestPingcapParse(t *testing.T) {
	var sqls = []string{
		// `select "1", col as c, * from tb as t where c1  = 2`,
		// `select "a", "1", 1, c1 as c, c2, (c2), t. c3, count(c4) as cnt, sum(c5+c6) as s, * from db.tb as t where c7 = 2 union select * from t2`,
		`select binary "a"`,
	}
	for _, sql := range sqls {
		stmt, err := PingcapParse(sql)
		if err != nil {
			t.Error(err.Error())
		}
		pretty.Println(stmt)
	}
}

func TestCockroachDBParse(t *testing.T) {
	var sqls = []string{
		`select col as c, sum(c1) as s from tb as t`,
	}
	for _, sql := range sqls {
		stmt, err := CockroachDBParse(sql)
		if err != nil {
			t.Error(err.Error())
		}
		pretty.Println(stmt)
	}
}

func TestMSSQLParse(t *testing.T) {
	var sqls = []string{
		"SELECT col as c, SUM(c1) as s FROM `tb` AS t",
		"select * from tb limit 1",
	}
	for _, sql := range sqls {
		_, err := MSSQLParse(sql)
		if err != nil {
			t.Error(err.Error())
		}
		// pretty.Println(stmts)
	}
}

func TestPTParse(t *testing.T) {
	var sqls = []string{
		"SELECT col as c, SUM(c1) as s FROM `tb` AS t",
		"select * from tb limit 1",
	}
	for _, sql := range sqls {
		_, err := PTParse(sql)
		if err != nil {
			t.Error(err.Error())
		}
		// pretty.Println(stmts)
	}
}
