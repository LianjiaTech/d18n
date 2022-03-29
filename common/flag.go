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
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
	"github.com/jessevdk/go-flags"
)

var sessionPassword = `this can't be a password, just for init`
var sessionDatabase = ""

// https://darjun.github.io/2020/01/10/godailylib/go-flags/

type Option struct {
	Verbose []bool       `short:"v" long:"verbose" required:"false" description:"verbose mode"`
	Help    func() error `long:"help" required:"false" description:"Show this help message"`

	// database config
	Server              string `long:"server" default:"mysql" description:"server type, support: mysql, postgres, sqlite, oracle, sqlserver, clickhouse"`
	DSN                 string `long:"dsn" description:"formatted data source name"`
	User                string `short:"u" long:"user" description:"database user"`
	Password            string `long:"password" description:"database password"`
	InteractivePassword bool   `short:"p" description:"input password interactively"`
	DefaultsExtraFile   string `long:"defaults-extra-file" description:"like mysql --defaults-extra-file for hidden password"`
	LoginPath           string `long:"login-path" description:"read config like mysql login-path"`
	Host                string `short:"h" long:"host" default:"127.0.0.1" description:"database host"`
	Port                int    `short:"P" long:"port" default:"3306" description:"database port"`
	Socket              string `short:"S" long:"socket" description:"unix socket file"`
	Database            string `short:"d" long:"database" description:"mysql/sql server: database name, oracle: service_name/sid, sqlite: database file path, csv: database directory"`
	Table               string `long:"table" description:"table name"`
	Charset             string `long:"charset" default:"utf8mb4" description:"connection charset"`
	Limit               int    `long:"limit" description:"query result lines limit"`
	Timeout             int    `long:"timeout" description:"query timeout in seconds"`

	// read from fille
	Query            string `short:"e" long:"query" description:"query read from file or command line"`
	Parser           string `long:"parser" default:"pingcap" description:"query parser: tidb, cockroach"`
	Vertical         bool   // print result vertical
	Prompt           string `long:"prompt" description:"iteractive query prompt"`
	InteractiveQuery bool   `short:"q" description:"input query interactively"`
	File             string `short:"f" long:"file" description:"input/output file"`
	Schema           string `long:"schema" description:"schema config file. support: sql, txt"`
	Mask             string `long:"mask" description:"data masking config file. support: csv, psv, tsv format"`
	Cipher           string `long:"cipher" description:"cipher config file. support: yaml"`
	Sensitive        string `long:"sensitive" description:"sensitive detection config file. support: yaml"`

	// call tools
	PrintCipher bool   `long:"print-cipher" description:"print or auto-generate cipher"`
	PrintConfig bool   `long:"print-config" description:"print config"`
	Preview     uint   `long:"preview" default:"0" description:"preview result file, print first N lines"`
	Lint        bool   `long:"lint" description:"lint file"`
	Emport      bool   `long:"import" description:"import file into database"`
	Detect      bool   `long:"detect" description:"detect sensitive info from data"`
	Watermark   string `long:"watermark" default:"" description:"watermark in export file. support: html, xlsx"`

	// sql config
	CheckEmpty              bool   `long:"check-empty" description:"check query result, if empty raise error"`
	Replace                 bool   `long:"replace" description:"generate sql use replace into syntax, only support MySQL and SQLite"`
	Update                  string `long:"update" default:"" description:"update primary key, separate by comma, case insensitive"`
	CompleteInsert          bool   `long:"complete-insert" description:"complete insert with columns name"`
	HexBLOB                 string `long:"hex-blob" description:"need hex encoding columns, separate by comma, case insensitive"`
	IgnoreColumns           string `long:"ignore-columns" default:"" description:"import file ignore columns, separated by comma"`
	ExtendedInsert          uint   `long:"extended-insert" default:"1" description:"use multiple-row INSERT syntax that include several values list"`
	ANSIQuotes              bool   `long:"ansi-quotes" description:"enable ANSI_QUOTES"`
	DisableForeignKeyChecks bool   `long:"disable-foreign-key-checks" description:"disable foreign key checks"`

	// other config
	BOM              bool   `long:"bom" description:"csv file with UTF8 BOM"`
	ExcelMaxFileSize uint   `long:"excel-max-file-size" description:"excel max file size, limit by memory"`
	LintLevel        string `long:"lint-level" default:"error" description:"file lint level"`
	IgnoreBlank      bool   `long:"ignore-blank" description:"ignore blank lines or columns when import file"`
	Comma            string `long:"comma" default:"," description:"csv comma char"`
	NoHeader         bool   `long:"no-header" description:"no header line, only data lines"`
	Comments         string `long:"comments" default:"#,--" description:"support comment characters, multiple comment split by comma"`
	SkipLines        uint   `long:"skip-lines" default:"0" description:"skip first N lines"`
	RandSeed         int64  `long:"rand-seed" description:"random seed, default: current unix nano timestamp"`
	MaxBufferSize    int    `long:"max-buffer-size" description:"bufio MaxScanTokenSize"`
	NULLString       string `long:"null-string" default:"NULL" description:"NULL string write into file. e.g., NULL, nil, None, \"\""`
}

// printPrompt print prompt
func (opt *Option) printPrompt() {
	var err error
	opt.Prompt, err = strconv.Unquote(`"` + opt.Prompt + `"`)
	if err == nil && opt.Prompt != "" {
		fmt.Print(opt.Prompt)
	} else {
		// allow line separator, sql end with ';'
		fmt.Print(opt.Server + " > ")
	}
}

// readQuery read query sql from stdin
func (opt *Option) readQuery() error {
	reader := bufio.NewReaderSize(os.Stdin, opt.MaxBufferSize)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		opt.Query = opt.Query + line
		line = strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(line), "use") ||
			strings.HasPrefix(strings.ToLower(line), "save") {
			break
		}
		switch strings.ToLower(strings.TrimSpace(opt.Query)) {
		case "exit", "quit":
			os.Exit(0)
		}

		if strings.HasSuffix(line, ";") || strings.HasSuffix(line, `\G`) {
			if strings.HasSuffix(line, `\G`) {
				opt.Vertical = true
				opt.Query = strings.TrimSuffix(strings.TrimSpace(opt.Query), `\G`)
			}
			break
		}
	}
	return nil
}

// prepareQuery deal query before send to database server
func (opt *Option) prepareQuery() error {
	var err error
	// use database
	dbReg := regexp.MustCompile(`(?i)^\s*use\s+[` + "`" + `\["']?(?P<Database>\w+)[` + "`" + `\]"']?\s*[;]?`)
	sub := dbReg.FindStringSubmatch(opt.Query)
	if len(sub) == 2 {
		sessionDatabase = sub[1]
		opt.Database = sessionDatabase
	}

	// interactive change save result type
	saveReg := regexp.MustCompile(`(?i)^\s*save\s+["']?(?P<File>\w+\.\w+)["']?\s*[;]?`)
	sub = saveReg.FindStringSubmatch(strings.TrimSpace(opt.Query))
	if len(sub) == 2 {
		opt.Query = ""
		opt.File = sub[1]
		opt.printPrompt()
		return opt.readQuery()
	}

	return err
}

func ParseFlags() (Config, error) {
	var err error
	var c Config
	var opt Option

	p := flags.NewParser(&opt, flags.Default & ^flags.HelpFlag)
	opt.Help = func() error {
		var b bytes.Buffer
		p.WriteHelp(&b)
		fmt.Println(b.String())
		os.Exit(0)
		return nil
	}
	p.Parse()

	if len(os.Args) == 1 {
		opt.Help()
		return c, fmt.Errorf("")
	}

	if opt.MaxBufferSize == 0 {
		opt.MaxBufferSize = bufio.MaxScanTokenSize
	}
	if opt.RandSeed == 0 {
		opt.RandSeed = time.Now().UnixNano()
	}
	if opt.ExcelMaxFileSize == 0 {
		opt.ExcelMaxFileSize = DefaultExcelMaxFileSize
	}

	if opt.DefaultsExtraFile != "" {
		err := parseDefaultsExtraFile(opt.DefaultsExtraFile, &c)
		if err != nil {
			return c, err
		}
	}

	if opt.LoginPath != "" {
		err := parseLoginPath(opt.LoginPath, &c)
		if err != nil {
			return c, err
		}
	}

	if c.User != "" {
		opt.User = c.User
	}
	if c.Password != "" {
		opt.Password = c.Password
	}
	if c.Charset != "" {
		opt.Charset = c.Charset
	}
	if c.Host != "" {
		opt.Host = c.Host
	}
	if c.Port != "" {
		opt.Port, err = strconv.Atoi(c.Port)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
	if c.Database != "" {
		opt.Database = c.Database
	} else if sessionDatabase != "" {
		opt.Database = sessionDatabase
	}

	if len(opt.Comma) > 1 || opt.Comma == "" {
		println("csv comma is rune, only one character")
		os.Exit(1)
	}

	if !strings.HasPrefix(strings.ToLower(opt.Charset), "utf") {
		opt.BOM = false
	}

	if opt.InteractivePassword {
		if sessionPassword != `this can't be a password, just for init` {
			opt.Password = sessionPassword
		} else {
			fmt.Print("Password: ")
			password, err := gopass.GetPasswd()
			if err != nil {
				return c, err
			}
			opt.Password = strings.TrimSpace(string(password))
			sessionPassword = opt.Password
		}
	}

	// read query interactive
	if opt.InteractiveQuery {
		// print prompt prefix
		opt.printPrompt()

		// read query interactive
		if err := opt.readQuery(); err != nil {
			return c, err
		}

		// prepare query before send to database server
		if err := opt.prepareQuery(); err != nil {
			return c, err
		}
	}

	// --query is empty and -table not empty generate full select query
	if opt.Table != "" && opt.Query == "" {
		if opt.Limit != 0 {
			c.Server = opt.Server
			switch strings.ToLower(opt.Server) {
			case "oracle":
				opt.Query = fmt.Sprintf("SELECT * FROM %s WHERE ROWNUM <= %d", c.QuoteKey(opt.Table), opt.Limit)
			case "sqlserver", "mssql":
				opt.Query = fmt.Sprintf("SELECT TOP %d * FROM %s", opt.Limit, c.QuoteKey(opt.Table))
			default:
				opt.Query = fmt.Sprintf("SELECT * FROM %s LIMIT %d", c.QuoteKey(opt.Table), opt.Limit)
			}
		} else {
			opt.Query = fmt.Sprintf("SELECT * FROM %s", c.QuoteKey(opt.Table))
		}
	}

	// test read from file
	if _, err := os.Stat(opt.Query); err == nil {
		q, err := ReadFileString(opt.Query)
		if err == nil {
			opt.Query = q
		}
	}

	// use abs path
	pwd, err := os.Getwd()
	if err != nil {
		return c, err
	}
	if !filepath.IsAbs(opt.File) &&
		(opt.File != "" && opt.File != "stdout") {
		opt.File = filepath.Join(pwd, opt.File)
	}

	c = Config{
		Server:   opt.Server,
		User:     opt.User,
		Password: opt.Password,
		Charset:  opt.Charset,
		DSN:      opt.DSN,

		Host:     opt.Host,
		Socket:   opt.Socket,
		Port:     fmt.Sprint(opt.Port),
		Database: opt.Database,
		Limit:    opt.Limit,
		Timeout:  opt.Timeout,

		Interactive:             opt.InteractiveQuery,
		Vertical:                opt.Vertical,
		Query:                   opt.Query,
		Parser:                  opt.Parser,
		File:                    opt.File,
		Schema:                  opt.Schema,
		Mask:                    opt.Mask,
		Cipher:                  opt.Cipher,
		Sensitive:               opt.Sensitive,
		Verbose:                 opt.Verbose,
		CheckEmpty:              opt.CheckEmpty,
		BOM:                     opt.BOM,
		ANSIQuotes:              opt.ANSIQuotes,
		DisableForeignKeyChecks: opt.DisableForeignKeyChecks,
		NULLString:              opt.NULLString,
		MaxBufferSize:           opt.MaxBufferSize,
		PrintCipher:             opt.PrintCipher,
		PrintConfig:             opt.PrintConfig,
		IgnoreBlank:             opt.IgnoreBlank,
		ExtendedInsert:          int(opt.ExtendedInsert),

		RandSeed: opt.RandSeed,

		Preview:   int(opt.Preview),
		Lint:      opt.Lint,
		LintLevel: opt.LintLevel,
		Import:    opt.Emport,
		Detect:    opt.Detect,
		Watermark: opt.Watermark,

		ExcelMaxFileSize: int(opt.ExcelMaxFileSize),

		Comma:     rune((opt.Comma)[0]),
		NoHeader:  opt.NoHeader,
		Comments:  strings.Split(opt.Comments, ","),
		SkipLines: int(opt.SkipLines),

		Replace:        opt.Replace,
		Update:         parseCommaFlag(opt.Update),
		Table:          opt.Table,
		CompleteInsert: opt.CompleteInsert,
		HexBLOB:        parseCommaFlag(opt.HexBLOB),
		IgnoreColumns:  parseCommaFlag(opt.IgnoreColumns),
	}

	// Fields alias map for data mask with column alias
	// ignore all errors
	fields, _ := c.ParseSelectFields()
	c.FieldsAliasMap = fieldsAliasMap(fields)

	// get table name from file prefix
	if c.Table == "" {
		c.Table = strings.Split(filepath.Base(c.File), ".")[0]
	}

	if (c.Server == "sqlite" || c.Server == "sqlite3") &&
		c.Database == "" && c.DSN == "" {
		println("sqlite should specified `--database DATA_FILE` arg")
		os.Exit(1)
	}

	return c, err
}
