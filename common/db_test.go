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
)

func TestQueryRows(t *testing.T) {
	orgCfg := TestConfig
	fmt.Println(TestConfig)
	sqls := [][]string{
		// wrong cases
		{
			`select * from not_exist_table`, // not exist table
			`xxxx`,                          // wrong syntax
			// If sql kill by `kill query` or by `max_execution_time` go-sql-driver/mysql does not raise error, so you should check empty result for the right result.
			// ERROR 3024 (HY000): Query execution was interrupted, maximum statement execution time exceeded
			`select /*+ MAX_EXECUTION_TIME(10) */ * from sakila.actor where sleep(1);`,
		},
		// right cases
		{
			`select * from sakila.actor where 1=2`,
		},
	}

	// wrong cases
	for _, sql := range sqls[0] {
		TestConfig.Query = sql
		rows, err := TestConfig.QueryRows()
		if rows != nil {
			rows.Next()
			err = rows.Err()
		}

		if err == nil {
			t.Error("show get error")
		}
	}

	// right cases
	for _, sql := range sqls[1] {
		TestConfig.Query = sql
		_, err := TestConfig.QueryRows()
		if err != nil {
			t.Error(err.Error())
		}
	}

	TestConfig = orgCfg
}

func TestExecResult(t *testing.T) {
	orgCfg := TestConfig

	sqls := [][]string{
		{
			`select * from not_exist_table`, // not exist table
			`xxxx`,                          // wrong syntax
			`select /*+ MAX_EXECUTION_TIME(10) */ * from sakila.actor where sleep(1);`, // ERROR 3024 (HY000): Query execution was interrupted, maximum statement execution time exceeded
		},
		{
			"do 1",
		},
	}

	// wrong cases
	for _, sql := range sqls[0] {
		TestConfig.Query = sql
		_, err := TestConfig.ExecResult()
		if err == nil {
			t.Error("show get error")
		}
	}

	// right cases
	for _, sql := range sqls[1] {
		TestConfig.Query = sql
		_, err := TestConfig.ExecResult()
		if err != nil {
			t.Error(err.Error())
		}
	}

	TestConfig = orgCfg
}

func ExampleDBParseNullString() {
	var columns = Row{"abc", "", "NULL", "null"}
	header := make([]HeaderColumn, len(columns))
	fmt.Println(TestConfig.DBParseNullString(header, columns))
	// Output:
	// [{abc true} { true} {NULL false} {null true}]
}

func TestSetForeignKeyChecks(t *testing.T) {
	orgCfg := TestConfig

	conn, err := TestConfig.NewConnection()
	if err != nil {
		t.Error(err.Error())
	}
	err = TestConfig.SetForeignKeyChecks(true, conn, "actor")
	if err != nil {
		t.Error(err.Error())
	}

	TestConfig = orgCfg
}

func TestNewConnection(t *testing.T) {
	orgCfg := TestConfig

	servers := []string{
		"mysql",
		"postgres",
		"oracle",
		"sqlserver",
		"sqlite",
		"sqlite3",
		"csvq",
		"csv",
		"clickhouse",
		"presto",
		"h2",
	}
	for _, s := range servers {
		TestConfig.Server = s
		conn, err := TestConfig.NewConnection()
		if err != nil {
			t.Error(err.Error())
		}
		fmt.Println(s, conn)
	}

	TestConfig = orgCfg
}

func ExampleEscape() {
	fmt.Println(Escape("abc"))
	fmt.Println(Escape("abc'"))
	fmt.Println(Escape(`abc"`))
	fmt.Println(Escape(`abc\`))
	fmt.Println(Escape(`abc中文def`))
	// Output:
	// abc
	// abc\'
	// abc\"
	// abc\\
	// abc中文def
}

func ExampleQuoteString() {
	orgCfg := TestConfig

	fmt.Println(TestConfig.QuoteString("abc"))
	fmt.Println(TestConfig.QuoteString(`abc"`))
	fmt.Println(TestConfig.QuoteString(`abc'`))
	TestConfig.Server = "oracle"
	fmt.Println(TestConfig.QuoteString("oracle"))
	fmt.Println(TestConfig.QuoteString(`abc"`))
	fmt.Println(TestConfig.QuoteString(`abc'`))
	TestConfig.Server = "postgres"
	fmt.Println(TestConfig.QuoteString("postgres"))
	fmt.Println(TestConfig.QuoteString(`abc"`))
	fmt.Println(TestConfig.QuoteString(`abc'`))
	// Output:
	// "abc"
	// "abc\""
	// "abc\'"
	// 'oracle'
	// 'abc"'
	// 'abc'''
	// 'postgres'
	// 'abc"'
	// 'abc'''

	TestConfig = orgCfg
}

func ExampleQuoteKey() {
	orgCfg := TestConfig

	fmt.Println(TestConfig.QuoteKey("abc"))
	fmt.Println(TestConfig.QuoteKey(`abc"`))
	fmt.Println(TestConfig.QuoteKey(`abc'`))
	TestConfig.Server = "oracle"
	fmt.Println(TestConfig.QuoteKey("abc"))
	// Output:
	// `abc`
	// `abc"`
	// `abc'`
	// "abc"

	TestConfig = orgCfg
}
