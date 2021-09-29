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
	"database/sql"
	"fmt"
)

func ExampleSQLInsertPrefix() {
	orgCfg := Cfg
	var header = Row{"a", "b", "c", "d"}
	Cfg.Table = "tab"
	fmt.Println(SQLInsertPrefix(header))
	Cfg.CompleteInsert = true
	fmt.Println(SQLInsertPrefix(header))
	// Output:
	// INSERT INTO `tab`  VALUES  <nil>
	// INSERT INTO `tab` (`a`, `b`, `c`, `d`) VALUES  <nil>
	Cfg = orgCfg
}

func ExampleSQLInsertValues() {
	columns := []sql.NullString{
		{String: "", Valid: false},
		{String: "", Valid: true},
		{String: "NULL", Valid: false},
		{String: "NULL", Valid: true},
		{String: "ABC", Valid: true},
		{String: "abc", Valid: true},
		{String: "1", Valid: true},
		{String: "1.0", Valid: true},
		{String: "0.1", Valid: true},
		{String: "-1", Valid: true},
		{String: "1e+2", Valid: true},
		{String: "1E+2", Valid: true},
	}

	header := []HeaderColumn{
		{ScanType: "string"},
		{ScanType: "string"},
		{ScanType: "string"},
		{ScanType: "string"},
		{ScanType: "string"},
		{ScanType: "string"},
		{ScanType: "int"},
		{ScanType: "int"},
		{ScanType: "int"},
		{ScanType: "int"},
		{ScanType: "int"},
		{ScanType: "int"},
	}
	fmt.Println(SQLInsertValues(header, columns))
	// Output:
	// (NULL, "", NULL, "NULL", "ABC", "abc", 1, 1.0, 0.1, -1, 1e+2, 1E+2) <nil>
}

func ExampleTableTemplate() {
	org := Cfg
	schemas := []string{
		TestPath + "/test/schema.txt",
		TestPath + "/test/mysql.schema.sql",
		TestPath + "/test/oracle.schema.sql",
		TestPath + "/test/postgres.schema.sql",
		TestPath + "/test/sqlite.schema.sql",
		TestPath + "/test/sqlserver.schema.sql",
	}
	for _, schema := range schemas {
		Cfg.Schema = schema
		fmt.Println(TableTemplate())
	}
	// Output:
	// [{actor_id  SMALLINT} {first_name  VARCHAR} {last_name  VARCHAR} {last_update  TIMESTAMP}] <nil>
	// [{actor_id  SMALLINT} {first_name  VARCHAR} {last_name  VARCHAR} {last_update  TIMESTAMP}] <nil>
	// [{actor_id  numeric} {first_name  VARCHAR} {last_name  VARCHAR} {last_update  DATE}] <nil>
	// [{actor_id  integer} {first_name  character} {last_name  character} {last_update  timestamp}] <nil>
	// [{actor_id  numeric} {first_name  VARCHAR} {last_name  VARCHAR} {last_update  TIMESTAMP}] <nil>
	// [{actor_id  int} {first_name  VARCHAR} {last_name  VARCHAR} {last_update  DATETIME}] <nil>
	Cfg = org
}

// This test case will cause actor.xlsx modify every time, event if no content change.
// If you need test SetXlsxWatermark use `make test-mysql`
// func ExampleSetXlsxWatermark() {
// 	fmt.Println(SetXlsxWatermark(TestPath+"/test/actor.xlsx", "watermark text"))
// 	// Output:
// 	// <nil>
// }

func ExampleGetXlsxWatermark() {
	fmt.Println(GetXlsxWatermark(TestPath + "/test/actor.xlsx"))
	// Output:
	// watermark text <nil>
}
