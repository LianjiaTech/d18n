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
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	// postgres
	_ "github.com/lib/pq"
	// sqlite pure go driver
	_ "modernc.org/sqlite"
	// oracle pure go driver
	_ "github.com/sijms/go-ora/v2"
	// MS SQL purge go driver
	_ "github.com/denisenkom/go-mssqldb"
	// clickhouse
	_ "github.com/ClickHouse/clickhouse-go"
	// presto
	_ "github.com/prestodb/presto-go-client/presto"
	// csvq
	_ "github.com/mithrandie/csvq-driver"
	// hive
	_ "github.com/taozle/go-hive-driver"
	// h2
	_ "github.com/jmrobles/h2go"
)

func (c Config) GetColumnTypes() ([]*sql.ColumnType, error) {
	db, err := c.NewConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var ctx context.Context
	ctx = context.Background()
	if c.Timeout > 0 {
		ctx, _ = context.WithTimeout(ctx, time.Duration(c.Timeout)*time.Second)
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s LIMIT 0", c.QuoteKey(c.Table)))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows.ColumnTypes()
}

// QueryRows mysql query get rows
func (c Config) QueryRows() (*sql.Rows, error) {
	db, err := c.NewConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var ctx context.Context
	ctx = context.Background()
	if c.Timeout > 0 {
		ctx, _ = context.WithTimeout(ctx, time.Duration(c.Timeout)*time.Second)
	}
	return db.QueryContext(ctx, strings.TrimRight(strings.TrimSpace(c.Query), ";"))
}

// ExecResult mysql query get result
func (c Config) ExecResult() (sql.Result, error) {
	db, err := c.NewConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var ctx context.Context
	ctx = context.Background()
	if c.Timeout > 0 {
		ctx, _ = context.WithTimeout(ctx, time.Duration(c.Timeout)*time.Second)
	}

	return db.ExecContext(ctx, strings.TrimRight(strings.TrimSpace(c.Query), ";"))
}

// newConnection init database connection
// Go 各种数据库连接字符串汇总 | 鸟窝
// https://colobu.com/2019/01/10/drivers-connection-string-in-Go/
func (c Config) NewConnection() (*sql.DB, error) {
	var dsn string
	switch c.Server {
	case "mysql":
		dsn = c.dsnMySQL()
	case "postgres":
		dsn = c.dsnPostgres()
	case "sqlite", "sqlite3", "csvq", "csv":
		dsn = c.dsnFile()
	case "oracle":
		dsn = c.dsnOracle()
	case "sqlserver", "mssql":
		dsn = c.dsnSQLServer()
	case "clickhouse":
		dsn = c.dsnClickHouse()
	case "presto":
		dsn = c.dsnPresto()
	case "hive":
		dsn = c.dsnHive()
	case "h2":
		dsn = c.dsnH2()
	}
	// --dsn flag highest level
	if c.DSN != "" {
		dsn = strings.TrimSpace(c.DSN)
	}

	switch c.Server {
	case "sqlite", "sqlite3":
		c.Server = "sqlite"
	case "csvq", "csv":
		c.Server = "csvq"
	}
	return sql.Open(c.Server, dsn)
}

// SetForeignKeyChecks
func (c Config) SetForeignKeyChecks(enable bool, conn *sql.DB, args ...string) error {
	var err error
	var sql string
	switch c.Server {
	case "sqlite", "sqlite3":
		sql = fmt.Sprintf("pragma foreign_keys %v;", enable)
	case "mysql":
		sql = fmt.Sprintf("SET FOREIGN_KEY_CHECKS = %v;", enable)
	case "csvq", "csv", "clickhouse", "presto":
		return fmt.Errorf("not support foreign key")
	case "postgres":
		if enable {
			sql = fmt.Sprintf("ALTER TABLE %s ENABLE TRIGGER ALL;", args[0])
		} else {
			sql = fmt.Sprintf("ALTER TABLE %s DISABLE TRIGGER ALL;", args[0])
		}
	case "oracle", "sqlserver", "mssql", "hive":
		// Notice: not suport tmp disable foreign key check by session
		return err
	}

	if c.DBAvailable(conn) {
		_, err = conn.Exec(sql)
	} else {
		fmt.Println(sql)
	}

	return err
}

// dsnMySQL concat mysql dsn string
func (c Config) dsnMySQL() string {
	var dsn string
	if c.Socket != "" {
		dsn = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=%s",
			c.User, c.Password, c.Socket, c.Database, c.Charset)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
			c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	}
	return dsn
}

// dsnPostgres concat postgres dsn string
func (c Config) dsnPostgres() string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.Database)
	return dsn
}

// dsnFile sqlite, csv database file string
func (c Config) dsnFile() string {
	pwd, _ := os.Getwd()
	if !filepath.IsAbs(c.Database) {
		c.Database = filepath.Join(pwd, c.Database)
	}

	return c.Database
}

// dsnOracle concat oracle dsn string
func (c Config) dsnOracle() string {
	return fmt.Sprintf("oracle://%s:%s@%s:%s/%s",
		c.User, c.Password, c.Host, c.Port, c.Database,
	)
}

// dsnSQLServer concat sqlserver dsn string
func (c Config) dsnSQLServer() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		c.User, c.Password, c.Host, c.Port, c.Database,
	)
}

// dsnClickHouse concat ClickHouse dsn string
func (c Config) dsnClickHouse() string {
	return fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s",
		c.Host, c.Port, c.User, c.Password, c.Database,
	)
}

// dsnPresto concat PrestoDB dsn string
func (c Config) dsnPresto() string {
	return fmt.Sprintf("http://%s@%s:%s", c.User, c.Host, c.Port)
}

// dsnHive concat Hive dsn string
func (c Config) dsnHive() string {
	return fmt.Sprintf("hive://%s:%s@%s:%s", c.User, c.Password, c.Host, c.Port)
}

// dsnH2 concat H2 dsn string
func (c Config) dsnH2() string {
	if c.User == "" {
		return fmt.Sprintf("h2://%s:%s/%s", c.Host, c.Port, c.Database)
	}
	return fmt.Sprintf("h2://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.Database)
}

// DBParseNullString convert []string to []sql.NullString
func (c Config) DBParseNullString(header []HeaderColumn, columns []string) []sql.NullString {
	values := make([]sql.NullString, len(columns))
	for i, col := range columns {
		if col == c.NULLString {
			values[i] = sql.NullString{String: col, Valid: false}
		} else {
			values[i] = sql.NullString{String: col, Valid: true}
		}
	}
	return values
}

// DBParserColumnNames convert *sql.ColumnType to column name string list
func (c Config) DBParserColumnNames(header []*sql.ColumnType) []string {
	var columns []string
	for _, h := range header {
		columns = append(columns, h.Name())
	}
	return columns
}

// DBParseHeaderColumn convert []HeaderColumn to column name string list
func (c Config) DBParseHeaderColumn(header []HeaderColumn) []string {
	var columns []string
	for _, h := range header {
		columns = append(columns, h.Name)
	}
	return columns
}

// DBParseColumnTypes convert *sql.ColumnType to self define HeaderColumn list
func (c Config) DBParseColumnTypes(header []*sql.ColumnType) []HeaderColumn {
	var headerColumns []HeaderColumn
	for _, h := range header {
		var scanType string
		if h.ScanType() != nil { // some database drive will not set ScanType, eg. sqlite
			scanType = h.ScanType().Name()
		}
		headerColumns = append(headerColumns, HeaderColumn{
			Name:         h.Name(),
			ScanType:     scanType,
			DatabaseType: h.DatabaseTypeName(),
		})
	}
	return headerColumns
}

// DBAvailable ...
func (c Config) DBAvailable(conn *sql.DB) bool {
	if conn == nil {
		return false
	}
	err := conn.Ping()
	return err == nil
}

func (c Config) QuoteString(str string) string {
	// How to escape special characters in Oracle SQL?
	// http://www.e2college.com/blogs/oracle/oracle_pl_sql_sql_queries/how_to_escape_special_characters_in_oracle_sql_.html

	switch c.Server {
	case "postgres", "oracle", "sqlserver", "mssql", "clickhouse", "presto":
		return "'" + strings.Replace(str, "'", "''", -1) + "'"
	default: // mysql, mariadb, tidb, sqlite, csvq, hive
		if c.ANSIQuotes {
			return `'` + Escape(str) + `'`
		} else {
			return `"` + Escape(str) + `"`
		}
	}
}

func (c Config) QuoteKey(str string) string {
	switch c.Server {
	case "postgres", "oracle", "sqlserver", "mssql", "clickhouse", "presto":
		return strconv.Quote(str)
	default:
		// MySQL
		// backtick (`) can be used to delimit identifiers whether or not ANSI_QUOTES
		// is enabled, but if ANSI_QUOTES is enabled, then "you cannot use double
		// quotation marks to quote literal strings, because it is interpreted as an identifier."
		return "`" + str + "`" // mysql, mariadb, tidb, sqlite, csvq, hive
	}
}

// HexBLOB ...
func (c Config) Hex(name string, value interface{}) (string, bool) {
	var hexed bool
	for _, k := range c.HexBLOB {
		if k == "*" || strings.EqualFold(k, name) {
			hexed = true
			return c.hex(fmt.Sprint(value)), hexed
		}
	}
	return fmt.Sprint(value), hexed
}

// hexBLOB hex-blob column
func (c Config) hex(value interface{}) string {
	var ret = fmt.Sprint(value)

	if len(c.HexBLOB) == 0 {
		return ret
	}

	switch c.Server {
	case "postgres":
		return `'\x` + hex.EncodeToString([]byte(ret)) + "'::bytea"
	case "sqlite", "sqlite3":
		return "X'" + hex.EncodeToString([]byte(ret)) + "'"
	case "oracle":
		// https://www.sqlines.com/oracle/datatypes/raw
		return c.QuoteString(strings.ToUpper(hex.EncodeToString([]byte(ret))))
	case "csvq", "csv", "clickhouse", "presto":
		// not support binary data
	// TODO: hive
	default: // mysql, mariadb, tidb, sql server
		return "0x" + hex.EncodeToString([]byte(ret))
	}

	return ret
}

func Escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
		case '\n': /* Must be escaped for logs */
			escape = 'n'
		case '\r':
			escape = 'r'
		case '\\':
			escape = '\\'
		case '\'':
			escape = '\''
		case '"': /* Better safe than sorry */
			escape = '"'
		case '\032': /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}
