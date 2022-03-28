package common

import (
	"regexp"
	"strings"

	// mysql, mariadb, tidb
	pingcap "github.com/pingcap/parser"
	pingcapAst "github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"

	// postgres
	cockroachdb "github.com/auxten/postgresql-parser/pkg/sql/parser"
)

type SelectTables struct {
	Tables []SelectTable
}

type SelectTable struct {
	Database string
	Table    string
}

type SelectField struct {
	Database string
	Table    string
	Name     string
	As       string
}

func fieldsAliasMap(fields []SelectField) map[string]string {
	alias := make(map[string]string)
	for _, f := range fields {
		if f.As != "" && f.Name != "" {
			alias[f.As] = f.Name
		}
	}
	return alias
}

// ParseSelectFields ...
func (c Config) ParseSelectFields() (fields []SelectField, err error) {
	switch c.Server {
	case "mysql":
		stmts, err := PingcapParse(c.Query)
		if err != nil {
			return fields, err
		}
		return PingcapSelectFields(stmts)
	case "postgres":
		stmts, err := CockroachDBParse(c.Query)
		if err != nil {
			return fields, err
		}
		return CockroachDBSelectFields(stmts)
	}
	return fields, err
}

func (c Config) ParseSelectTables() (tables SelectTables, err error) {
	switch c.Server {
	case "mysql":
		stmts, err := PingcapParse(c.Query)
		if err != nil {
			return tables, err
		}
		v := &tables
		for _, stmt := range stmts {
			stmt.Accept(v)
		}
	case "postgres":
	}
	return tables, err
}

// PingcapParse use tidb parser for tidb/mysql/mariadb
func PingcapParse(sql string) ([]pingcapAst.StmtNode, error) {
	var charset, collation string
	p := pingcap.New()
	sql = removeIncompatibleWords(sql)
	stmt, _, err := p.Parse(sql, charset, collation)
	if err != nil {
		return stmt, err
	}
	return stmt, err
}

// removeIncompatibleWords remove pingcap/parser not support words from schema
func removeIncompatibleWords(sql string) string {
	fields := strings.Fields(strings.TrimSpace(sql))
	if len(fields) == 0 {
		return sql
	}
	switch strings.ToLower(fields[0]) {
	case "create", "alter":
	default:
		return sql
	}
	// CONSTRAINT col_fk FOREIGN KEY (col) REFERENCES tb (id) ON UPDATE CASCADE
	re := regexp.MustCompile(`(?i) ON UPDATE CASCADE`)
	sql = re.ReplaceAllString(sql, "")

	// FULLTEXT KEY col_fk (col) /*!50100 WITH PARSER `ngram` */
	// /*!50100 PARTITION BY LIST (col)
	re = regexp.MustCompile(`/\*!5`)
	sql = re.ReplaceAllString(sql, "/* 5")

	// col varchar(10) CHARACTER SET gbk DEFAULT NULL
	re = regexp.MustCompile(`(?i)CHARACTER SET [a-z_0-9]* `)
	sql = re.ReplaceAllString(sql, "")

	// CREATE TEMPORARY TABLE IF NOT EXISTS t_film AS (SELECT * FROM film);
	re = regexp.MustCompile(`(?i)CREATE TEMPORARY TABLE`)
	sql = re.ReplaceAllString(sql, "CREATE TABLE")

	return sql
}

// PingcapSelectFields only get output fields not all columns in select clause
func PingcapSelectFields(stmts []pingcapAst.StmtNode) ([]SelectField, error) {
	var fields []SelectField
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *pingcapAst.SelectStmt:
			for _, f := range s.Fields.Fields {
				if f.WildCard == nil {
					switch expr := f.Expr.(type) {
					case *pingcapAst.ColumnNameExpr:
						fields = append(fields, SelectField{
							Database: expr.Name.Schema.L,
							Table:    expr.Name.Table.L,
							Name:     expr.Name.Name.L,
							As:       f.AsName.L,
						})
					}
				} else {
					fields = append(fields, SelectField{
						Name: "*",
					})
				}
			}
		}
	}
	return fields, nil
}

// https://github.com/pingcap/parser/blob/master/docs/quickstart.md#traverse-ast-nodes

// Enter tables ast node visitor
func (v *SelectTables) Enter(in pingcapAst.Node) (pingcapAst.Node, bool) {
	if name, ok := in.(*pingcapAst.TableName); ok {
		v.Tables = append(v.Tables, SelectTable{Database: name.Schema.L, Table: name.Name.L})
	}
	return in, false
}

// Leave tables ast node visitor
func (v *SelectTables) Leave(in pingcapAst.Node) (pingcapAst.Node, bool) {
	return in, true
}

// CockroachDBParse use cockroachdb parser for postgres/cockroachdb
func CockroachDBParse(sql string) (cockroachdb.Statements, error) {
	return cockroachdb.Parse(sql)
}

func CockroachDBSelectFields(stmts cockroachdb.Statements) (fields []SelectField, err error) {

	return fields, err
}
