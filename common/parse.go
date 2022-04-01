package common

import (
	"regexp"
	"strconv"
	"strings"

	// mysql, mariadb, tidb
	pingcap "github.com/pingcap/parser"
	pingcapAst "github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"

	// postgres, cockroach
	cockroachDB "github.com/auxten/postgresql-parser/pkg/sql/parser"
	cockroachTree "github.com/auxten/postgresql-parser/pkg/sql/sem/tree"
	cockroachWalk "github.com/auxten/postgresql-parser/pkg/walk"
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

type SelectFields struct {
	Fields []SelectField
}

type SelectFuncs struct {
	Funcs []string
}

func fieldsAliasMap(fields SelectFields) map[string]string {
	alias := make(map[string]string)
	for _, f := range fields.Fields {
		if f.As != "" && f.Name != "" {
			alias[f.As] = f.Name
		}
	}
	return alias
}

// ParseSelectFields ...
func (c Config) ParseSelectFields() (fields SelectFields, err error) {
	switch strings.ToLower(c.Parser) {
	case "mysql", "pingcap", "tidb", "mariadb", "":
		stmts, err := PingcapParse(c.Query)
		if err != nil {
			return fields, err
		}
		return PingcapSelectFields(stmts)
	case "postgres", "cockroachdb", "cockroach":
		stmts, err := CockroachDBParse(c.Query)
		if err != nil {
			return fields, err
		}
		return CockroachDBSelectFields(stmts)
	}
	return fields, err
}

func (c Config) ParseSelectTables() (tables SelectTables, err error) {
	switch strings.ToLower(c.Parser) {
	case "mysql", "pingcap", "tidb", "mariadb", "":
		stmts, err := PingcapParse(c.Query)
		if err != nil {
			return tables, err
		}
		v := &tables
		for _, stmt := range stmts {
			stmt.Accept(v)
		}
	case "postgres", "cockroachdb", "cockroach":
		stmts, err := CockroachDBParse(c.Query)
		if err != nil {
			return tables, err
		}
		return CockroachDBSelectTables(stmts)
	}
	return tables, err
}

func (c Config) ParseSelectFuncs() (funcs SelectFuncs, err error) {
	switch strings.ToLower(c.Parser) {
	case "mysql", "pingcap", "tidb", "mariadb", "":
		stmts, err := PingcapParse(c.Query)
		if err != nil {
			return funcs, err
		}
		v := &funcs
		for _, stmt := range stmts {
			stmt.Accept(v)
		}
	case "postgres", "cockroachdb", "cockroach":
		stmts, err := CockroachDBParse(c.Query)
		if err != nil {
			return funcs, err
		}
		return CockroachDBSelectFuncs(stmts)
	}
	return funcs, err
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
func PingcapSelectFields(stmts []pingcapAst.StmtNode) (fields SelectFields, err error) {
	for _, stmt := range stmts {
		stmt.Accept(&fields)
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

// Enter tables ast node visitor
func (v *SelectFields) Enter(in pingcapAst.Node) (pingcapAst.Node, bool) {
	if f, ok := in.(*pingcapAst.SelectField); ok {
		var db, tb, col, as string
		// as
		as = f.AsName.L

		// db, tb, col
		val, err := strconv.Unquote(strings.ToLower(f.Text()))
		if err != nil {
			val = f.Text()
		}
		if val != "" {
			val = strings.TrimSpace(strings.TrimRight(strings.TrimLeft(val, "("), ")"))
			col = val
		}
		if f.WildCard == nil {
			switch c := f.Expr.(type) {
			case *pingcapAst.ColumnNameExpr:
				db = c.Name.Schema.L
				tb = c.Name.Table.L
				col = c.Name.Name.L
			}
		} else {
			col = "*"
		}

		if col != "" {
			v.Fields = append(v.Fields, SelectField{Database: db, Table: tb, Name: col, As: as})
		}
	}

	return in, false
}

// Leave tables ast node visitor
func (v *SelectFields) Leave(in pingcapAst.Node) (pingcapAst.Node, bool) {
	return in, true
}

// Enter tables ast node visitor
func (v *SelectFuncs) Enter(in pingcapAst.Node) (pingcapAst.Node, bool) {
	switch f := in.(type) {
	case *pingcapAst.AggregateFuncExpr:
		v.Funcs = append(v.Funcs, strings.ToLower(f.F))
	case *pingcapAst.FuncCallExpr:
		v.Funcs = append(v.Funcs, f.FnName.L)
	case *pingcapAst.WindowFuncExpr:
		v.Funcs = append(v.Funcs, strings.ToLower(f.F))
	case *pingcapAst.FuncCastExpr:
		var name string
		switch f.FunctionType {
		case 1:
			name = "cast"
		case 2:
			name = "convert"
		case 3:
			name = "binary"
		}
		v.Funcs = append(v.Funcs, name)
	}
	return in, false
}

// Leave tables ast node visitor
func (v *SelectFuncs) Leave(in pingcapAst.Node) (pingcapAst.Node, bool) {
	return in, true
}

// CockroachDBParse use CockroachDB parser for postgres/cockroachdb
func CockroachDBParse(sql string) (cockroachDB.Statements, error) {
	return cockroachDB.Parse(sql)
}

func CockroachDBSelectFields(stmts cockroachDB.Statements) (fields SelectFields, err error) {
	w := &cockroachWalk.AstWalker{
		UnknownNodes: []interface{}{},
		Fn: func(ctx interface{}, node interface{}) (stop bool) {
			if n, ok := node.(cockroachTree.SelectExpr); ok {
				var db, tb, col, as string
				// as
				as, err = strconv.Unquote(strings.ToLower(n.As.String()))
				if err != nil {
					as = n.As.String()
				}

				// db, tb, col
				tup := strings.Split(strings.ToLower(n.Expr.String()), ".")
				switch len(tup) {
				case 1:
					col = tup[0]
				case 2:
					tb = tup[0]
					col = tup[1]
				case 3:
					db = tup[0]
					tb = tup[1]
					col = tup[2]
				}
				val, _ := strconv.Unquote(strings.ToLower(col))
				if val != "" {
					col = val
				}
				if strings.HasPrefix(col, "(") {
					col = strings.TrimSpace(strings.TrimRight(strings.TrimLeft(col, "("), ")"))
				}
				if strings.HasSuffix(col, ")") {
					col = ""
				}

				if col != "" {
					fields.Fields = append(fields.Fields, SelectField{Database: db, Table: tb, Name: col, As: as})
				}
			}
			return false
		},
	}

	_, err = w.Walk(stmts, nil)
	return fields, err
}

func CockroachDBSelectTables(stmts cockroachDB.Statements) (tables SelectTables, err error) {
	w := &cockroachWalk.AstWalker{
		Fn: func(ctx interface{}, node interface{}) (stop bool) {
			if n, ok := node.(*cockroachTree.AliasedTableExpr); ok {
				tables.Tables = append(tables.Tables,
					SelectTable{
						Database: strings.ToLower(n.Expr.(*cockroachTree.TableName).Schema()),
						Table:    strings.ToLower(n.Expr.(*cockroachTree.TableName).Table()),
					})
			}
			return false
		},
	}
	_, err = w.Walk(stmts, nil)
	return tables, err
}

func CockroachDBSelectFuncs(stmts cockroachDB.Statements) (funcs SelectFuncs, err error) {
	w := &cockroachWalk.AstWalker{
		Fn: func(ctx interface{}, node interface{}) (stop bool) {
			if f, ok := node.(*cockroachTree.FuncExpr); ok {
				funcs.Funcs = append(funcs.Funcs, f.Func.String())
			}
			return false
		},
	}
	_, err = w.Walk(stmts, nil)
	return funcs, err
}
