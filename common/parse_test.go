package common

import (
	"testing"

	"github.com/kr/pretty"
)

func ExampleParseSelectFields() {
	orgCfg := TestConfig
	TestConfig.Parser = "mysql"
	TestConfig.Query = `select "a", "1", 1, c1 as c, c2, t. c3, * from db.tb as t`
	pretty.Println(TestConfig.ParseSelectFields())
	TestConfig.Parser = "postgres"
	TestConfig.Query = `select "a", "1", 1, c1 as c, c2, t. c3, * from db.tb as t`
	pretty.Println(TestConfig.ParseSelectFields())
	// Output:
	// []common.SelectField{
	//     {Database:"", Table:"", Name:"c1", As:"c"},
	//     {Database:"", Table:"", Name:"c2", As:""},
	//     {Database:"", Table:"t", Name:"c3", As:""},
	//     {Database:"", Table:"", Name:"*", As:""},
	// } nil
	// []common.SelectField{
	//     {Database:"", Table:"", Name:"a", As:""},
	//     {Database:"", Table:"", Name:"\"1\"", As:""},
	//     {Database:"", Table:"", Name:"1", As:""},
	//     {Database:"", Table:"", Name:"c2", As:""},
	//     {Database:"", Table:"t", Name:"c3", As:""},
	//     {Database:"", Table:"", Name:"*", As:""},
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

func TestPingcapParse(t *testing.T) {
	var sqls = []string{
		`select col as c from tb as t`,
	}
	for _, sql := range sqls {
		_, err := PingcapParse(sql)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func TestCockroachDBParse(t *testing.T) {
	var sqls = []string{
		`select col as c from tb as t`,
	}
	for _, sql := range sqls {
		_, err := CockroachDBParse(sql)
		if err != nil {
			t.Error(err.Error())
		}
	}
}
