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
	"strings"

	"github.com/dolmen-go/mylogin"
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
	Timeout  int // query timeout

	// other config
	Interactive             bool   // interactive mode
	Vertical                bool   // print result vertical like MySQL '\G'
	Query                   string // select query
	File                    string // storage file abs path
	Schema                  string // create table sql file, use for import data
	Cipher                  string // cipher config file
	Mask                    string // mask config file, csv format
	Sensitive               string // sensitive data detection config
	Verbose                 []bool // verbose mod
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
func parseDefaultsExtraFile(file string, c *Config) error {
	config, err := ini.Load(file)
	if err != nil {
		return err
	}

	// get config from [client] section
	c.User = config.Section("client").Key("user").String()
	c.Password = config.Section("client").Key("password").String()
	c.Database = config.Section("client").Key("database").String()
	c.Host = config.Section("client").Key("host").String()
	c.Port = config.Section("client").Key("port").String()
	c.Charset = config.Section("client").Key("default-character-set").String()

	return err
}

// parseLoginPath parse --login-path file
func parseLoginPath(section string, c *Config) error {
	if section == "" {
		section = "client"
	}

	login, err := mylogin.ReadLogin(mylogin.DefaultFile(), []string{section})

	if login == nil {
		return fmt.Errorf("--login-path has not such section '%s'", section)
	}

	if login.User != nil {
		c.User = *login.User
	}
	if login.Password != nil {
		c.Password = *login.Password
	}
	if login.Host != nil {
		c.Host = *login.Host
	}
	if login.Port != nil {
		c.Port = *login.Port
	}
	return err
}

func (c Config) ParseSchema() (header []HeaderColumn, err error) {
	if c.Schema != "" {
		// 1. TableTemplate
		header, err = c.TableTemplate()
		if err != nil {
			return
		}
	} else {
		// 2. GetColumnTypes -> DBParserColumnNames
		var columns []*sql.ColumnType
		columns, err = c.GetColumnTypes()
		if err != nil {
			return
		}
		header = c.DBParseColumnTypes(columns)
	}
	return
}

func PrintConfig(c Config) {
	buf, err := yaml.Marshal(c)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(string(buf))
}
