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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	xlsx "github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Rows ...
type Rows []Row

// Rows ...
type Row []string

// HeaderColumn ...
type HeaderColumn struct {
	Name         string
	ScanType     string
	DatabaseType string
}

/*
# INSERT SYNTAX

Oracle:
https://docs.oracle.com/cd/B14117_01/appdev.101/b10807/13_elems025.htm

SQL Server:
https://docs.microsoft.com/en-us/sql/t-sql/statements/insert-transact-sql?view=sql-server-ver15

SQLite:
https://www.sqlite.org/lang_insert.html

MySQL:
https://dev.mysql.com/doc/refman/8.0/en/insert.html

PostgreSQL:
https://www.postgresql.org/docs/9.5/sql-insert.html

ClickHouse:
https://clickhouse.tech/docs/en/sql-reference/statements/insert-into/#insert

# REPLACE SYNTAX

MySQL:
https://dev.mysql.com/doc/refman/8.0/en/replace.html

SQLite:
https://www.sqlitetutorial.net/sqlite-replace-statement/

Notice:
* table name with backtick quote, mysql backtick, oracle double quote
* value with double quote, single quote
* not support replace into: presto, postgresql, oracle, clickhouse, sql server

*/

// SQLInsertPrefix ...
func (c Config) SQLInsertPrefix(header Row) (string, error) {
	var prefix string
	var err error

	if c.Table == "" {
		return prefix, fmt.Errorf("no table name")
	}

	// complete insert statement
	var columnName string
	var columnFilter []string
	if c.CompleteInsert {
		for i, v := range header {
			var ignore bool
			for _, name := range c.IgnoreColumns {
				if strings.EqualFold(name, v) {
					ignore = true
					break
				}
			}
			header[i] = c.QuoteKey(v)
			if !ignore {
				columnFilter = append(columnFilter, header[i])
			}
		}
		columnName = fmt.Sprintf("(%s)", strings.Join(columnFilter, ", "))
	}

	// Table name
	tableName := fmt.Sprint(c.QuoteKey(c.Table))

	// INSERT
	prefix = fmt.Sprintf("INSERT INTO %s %s VALUES ",
		tableName, columnName)

	// REPLACE
	if c.Replace {
		prefix = fmt.Sprintf("REPLACE INTO %s %s VALUES ",
			tableName, columnName)
	}

	// UPDATE
	if len(c.Update) > 0 {
		prefix = fmt.Sprintf("UPDATE %s SET ", tableName)
	}

	return prefix, err
}

// SQLInsertValues concat values to string
func (c Config) SQLInsertValues(header []HeaderColumn, columns []sql.NullString) (string, error) {
	if len(header) != len(columns) {
		return "", fmt.Errorf("%s, header: %v", WrongArgsCount, header)
	}

	var values []string
	var updateWhere []string
	var updateSet []string
	for i, col := range columns {
		// --ignore-columns
		var ignore bool
		for _, name := range c.IgnoreColumns {
			if strings.EqualFold(name, header[i].Name) {
				ignore = true
				break
			}
		}
		if ignore {
			continue
		}

		var value string
		if !col.Valid {
			value = c.NULLString
			// UPDATE set
			if len(c.Update) > 0 {
				updateSet = append(updateSet, fmt.Sprintf("%s=NULL", c.QuoteKey(header[i].Name)))
			}

			// UPDATE where
			for _, k := range c.Update {
				if strings.EqualFold(k, header[i].Name) {
					updateWhere = append(updateWhere, fmt.Sprintf("%s IS NULL", c.QuoteKey(header[i].Name)))
					break
				}
			}
		} else {
			switch header[i].ScanType {
			case "int", "int8", "int16", "int32", "int64",
				"unit", "uint8", "uint16", "uint32", "uint64",
				"float32", "float64":
				value = col.String

			default:
				// header[i].DatabaseTypeName() is not cross database compatible
				// float value like 0.1 ScanType is RawBytes, DatabaseTypeName is DECIMAL
				ty := strings.ToUpper(header[i].DatabaseType)
				switch ty {
				case "DECIMAL", "INT", "BIGINT", "INTEGER", "FLOAT", "DOUBLE", "SMALLINT",
					"NUMBER", "NUMERIC", "REAL", "SINGLE", "CURRENCY", "AUTONUMBER", "SMALLMONEY", "MONEY":
					value = col.String
				case "RAW": // Oracle RAW data
					value = c.QuoteString(strings.ToUpper(hex.EncodeToString([]byte(col.String))))
				default:
					// LONG: in Oracle means string, in SQL Server means big integer
					if c.Server == "sqlserver" && strings.EqualFold(header[i].DatabaseType, "LONG") {
						value = col.String
						break
					}

					// hex-blob
					v, hexed := c.Hex(header[i].Name, col.String)
					if hexed {
						value = v
					} else {
						switch ty {
						case "NVARCHAR", "NCHAR", "NATIONAL", "NVARCHAR2":
							value = "N" + c.QuoteString(col.String)
						default:
							value = c.QuoteString(col.String)
						}
					}
				}
			}

			// UPDATE set
			if len(c.Update) > 0 {
				updateSet = append(updateSet, fmt.Sprintf("%s=%s", c.QuoteKey(header[i].Name), value))
			}

			// UPDATE where
			for _, k := range c.Update {
				if strings.EqualFold(k, header[i].Name) {
					updateWhere = append(updateWhere, fmt.Sprintf("%s=%s", c.QuoteKey(header[i].Name), value))
					break
				}
			}
		}

		values = append(values, value)
	}

	if len(c.Update) > 0 && len(updateWhere) > 0 {
		return fmt.Sprintf("%s WHERE %s",
			strings.Join(updateSet, ", "),
			strings.Join(updateWhere, " AND "),
		), nil
	}

	return fmt.Sprintf("(%s)", strings.Join(values, ", ")), nil
}

func (c Config) SQLMultiValues(counter int, prefix, values string) string {
	delimiter := ";\n"

	// skip-extended-insert
	if c.ExtendedInsert <= 1 {
		return fmt.Sprint(prefix, values, delimiter)
	}

	// UPDATE
	if len(c.Update) > 0 {
		return fmt.Sprint(prefix, values, delimiter)
	}

	// INSERT, REPLACE
	switch counter % c.ExtendedInsert {
	case 1:
		delimiter = ""
	case 0:
		prefix = ", "
		delimiter = ";\n"
	default:
		prefix = ", "
		delimiter = ""
	}

	return fmt.Sprint(prefix, values, delimiter)
}

// TableTemplate get header []HeaderColumn from table tamplate
// schema must be strict formatted, only support SHOW CREATE output info
func (c Config) TableTemplate() ([]HeaderColumn, error) {
	var header []HeaderColumn

	createTable, err := ReadFileString(c.Schema)
	if err != nil {
		return header, err
	}

	columns := strings.Split(createTable, "\n")
	for _, col := range columns {
		col = strings.TrimSpace(col)
		def := strings.Fields(col)
		if len(def) >= 2 {
			// table name
			if def[0] == "CREATE" && def[1] == "TABLE" {
				if c.Table == "" && len(def) > 2 {
					c.Table = strings.Trim(def[2], "`\"'")
				}
				continue
			}

			// ignore comment line
			var ignore bool
			for _, comment := range c.Comments {
				if strings.HasPrefix(def[0], comment) {
					ignore = true
					break
				}
			}
			if ignore {
				continue
			}

			// keys
			if def[0] == "PRIMARY" && def[1] == "KEY" {
				break
			}
			if def[0] == "UNIQUE" && def[1] == "KEY" {
				break
			}
			if def[0] == "KEY" {
				break
			}
			if def[0] == "CONSTRAINT" {
				break
			}

			// table options
			if def[0] == ")" {
				break
			}

			// columns
			header = append(header, HeaderColumn{
				Name:         strings.Trim(def[0], "`\"'"),
				DatabaseType: strings.Split(def[1], "(")[0],
			})
		}
	}
	return header, err
}

// closedSemicolon quote close and line break
func closedSemicolon(data []byte) (int, error) {
	var quoteClosed = true
	var quote byte
	var commentClosed = true // like /* */
	var comment string
	for head := 0; head < len(data); head++ {
		if commentClosed && quoteClosed {
			if data[head] == '/' && data[head+1] == '*' {
				commentClosed = false
				comment = "/*"
			}
			if data[head] == '#' {
				commentClosed = false
				comment = "#"
			}
			if data[head] == '-' && data[head+1] == '-' && data[head+2] == ' ' {
				commentClosed = false
				comment = "-- "
			}
			if data[head] == '"' {
				quote = '"'
				quoteClosed = false
			} else if quoteClosed && data[head] == '\'' {
				quote = '\''
				quoteClosed = false
			}
		} else {
			if data[head] == quote {
				quoteClosed = true
			}
			if data[head] == '*' && data[head+1] == '/' && comment == "/*" {
				commentClosed = true
			}
			if data[head] == '\n' && (comment == "#" || comment == "-- ") {
				commentClosed = true
			}
		}
		if commentClosed && quoteClosed && data[head] == ';' {
			return head, nil
		}
	}
	return -1, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// SQLReadLine bufio.Scan() SplitFunc
func SQLReadLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return
	}

	// Read SQL format row from file
	if i, e := closedSemicolon(data); i > 0 {
		advance = i + 1
		token = data[0 : i+1]
		err = e
		return
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		advance = len(data)
		token = dropCR(data)
		return
	}

	// Request more data.
	return
}

func (c Config) ParseArray(values []string) string {
	var buf []byte
	var err error
	switch c.Server {
	case "clickhouse":
		for i, v := range values {
			values[i] = c.QuoteString(v)
		}
		buf = []byte("[" + strings.Join(values, ",") + "]")
	case "postgres":
		buf, err = json.Marshal(values)
		if err != nil {
			return fmt.Sprint(values)
		}
		if len(buf) >= 2 {
			buf[0] = '{'
			buf[len(buf)-1] = '}'
		}
	}
	return string(buf)
}

func SetXlsxWatermark(filename string, watermark string) error {
	fd, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}
	err = fd.SetDocProps(&xlsx.DocProperties{Title: watermark})
	if err != nil {
		return err
	}
	return fd.Save()
}

func GetXlsxWatermark(filename string) (string, error) {
	var watermark string
	fd, err := xlsx.OpenFile(filename)
	if err != nil {
		return watermark, err
	}
	attr, err := fd.GetDocProps()
	if err != nil {
		return watermark, err
	}
	watermark = attr.Title
	return watermark, err
}

// ReadFileString read string from file
func ReadFileString(filename string) (string, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// UTF8 BOM compatible
	if strings.HasPrefix(string(buf), UTF8BOM) {
		return strings.TrimPrefix(string(buf), UTF8BOM), err
	}
	return string(buf), err
}
