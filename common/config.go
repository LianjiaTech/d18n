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
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	ini "gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

// Config d18n config
type Config struct {
	// server config in my.cnf file, [client] section
	Server   string // mysql, postgres, oracle
	User     string
	Password string
	Charset  string
	DSN      string // Formated DSN string

	// other server config
	Host     string
	Socket   string
	Port     string
	Database string
	Limit    int // result lines limit

	// other config
	Query                   string // select query
	File                    string // storage file abs path
	Schema                  string // create table sql file, use for import data
	Cipher                  string // cipher config file
	Mask                    string // mask config file, csv format
	Sensitive               string // sensitive data detection config
	Verbose                 bool   // verbose mod
	CheckEmpty              bool   // check empty result, if empty raise error
	BOM                     bool   // add BOM ahead of plain text file, windows unicode chaos
	NULLString              string // NULL value write in file, e.g., NULL, None, nil, ""
	ANSIQuotes              bool   // enable ANSIQuotes
	DisableForeignKeyChecks bool   // disable foreign key checks
	MaxBufferSize           int    // bufio default buffer size
	PrintCipher             bool   // print d18n automatically generated cipher
	PrintConfig             bool   // print d18n config
	IgnoreBlank             bool   // ignore blank lines in import file
	ExtendedInsert          int    // mysqldump extended-insert

	// mask
	RandSeed int64 // rand.Seed()

	// tools
	Preview   int    // preview xlsx file, print first N lines
	Lint      bool   // file format check
	LintLevel string // lint break level
	Import    bool   // import file data into database
	Detect    bool   // detect sensitive info from data

	// digital watermark
	Watermark string // add watermark into html、xlsx
	// xlsx, sql, html

	// csv config
	Comma     rune     // Comma is the cell delimiter.
	Comments  []string // Comments are the line comment flags, e.g., #, --, //
	NoHeader  bool     // CSV file without Header line
	LineBreak string   // LineBreak whether it meets expectations
	SkipLines int      // skip first N lines when import data

	// excel config
	ExcelMaxFileSize int // excel file max size

	// sql config
	Replace        bool     // use replace into, instead of insert
	Update         []string // use update, instead of insert, primary key list, separated by comma
	Table          string   // table name
	CompleteInsert bool     // complete-insert
	HexBLOB        []string // blob column names
	IgnoreColumns  []string // ignore column list
}

// Cfg global config
var Cfg Config

func parseCommaFlag(update string) []string {
	var primary []string
	update = strings.TrimSpace(update)
	flags := strings.Split(update, ",")
	for _, f := range flags {
		f = strings.TrimSpace(f)
		if f != "" {
			primary = append(primary, f)
		}
	}
	return primary
}

// parseDefaultsExtraFile parse --defaults-extra-file file
func parseDefaultsExtraFile(file string) error {
	c, err := ini.Load(file)
	if err != nil {
		return err
	}

	// get config from [client] section
	Cfg.User = c.Section("client").Key("user").String()
	Cfg.Password = c.Section("client").Key("password").String()
	Cfg.Database = c.Section("client").Key("database").String()
	Cfg.Host = c.Section("client").Key("host").String()
	Cfg.Port = c.Section("client").Key("port").String()
	Cfg.Charset = c.Section("client").Key("default-character-set").String()

	return err
}

type MaskRule struct {
	MaskFunc string   `yaml:"func"`
	Args     []string `yaml:"args"`
}

var MaskConfig = make(map[string]MaskRule)

func ParseMaskConfig() error {

	// not config mask
	if Cfg.Mask == "" {
		return nil
	}

	fd, err := os.Open(Cfg.Mask)
	if err != nil {
		return err
	}
	defer fd.Close()

	r := csv.NewReader(fd)
	r.FieldsPerRecord = -1 // fix wrong number of fields
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(Cfg.Mask), "."))
	switch suffix {
	case "csv":
		r.Comma = ','
	case "psv":
		r.Comma = '|'
	case "tsv":
		r.Comma = '\t'
	case "txt":
		r.Comma = ' '
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}
	for {
		row, err := r.Read()
		if err == io.EOF { // end of file
			break
		} else if err != nil {
			return err
		}

		if len(row) > 1 {
			MaskConfig[strings.ToLower(row[0])] = MaskRule{
				MaskFunc: strings.ToLower(row[1]),
				Args:     row[2:],
			}
		}
	}
	return err
}

func ParseSchema() (header []HeaderColumn, err error) {
	if Cfg.Schema != "" {
		// 1. TableTemplate
		header, err = TableTemplate()
		if err != nil {
			return
		}
	} else {
		// 2. GetColumnTypes -> DBParserColumnNames
		var columns []*sql.ColumnType
		columns, err = GetColumnTypes()
		if err != nil {
			return
		}
		header = DBParseColumnTypes(columns)
	}
	return
}

func PrintConfig() {
	buf, err := yaml.Marshal(Cfg)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(string(buf))
}
