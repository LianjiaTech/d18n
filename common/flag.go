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
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
	"github.com/jessevdk/go-flags"
)

// https://darjun.github.io/2020/01/10/godailylib/go-flags/

func ParseFlags() error {
	var err error

	type option struct {
		Verbose bool         `short:"v" long:"verbose" required:"false" description:"verbose mode"`
		Help    func() error `long:"help" required:"false" description:"Show this help message"`

		// database config
		Server              string `long:"server" default:"mysql" description:"server type, support: mysql, postgres, sqlite, oracle, sqlserver, clickhouse"`
		DSN                 string `long:"dsn" description:"formatted data source name"`
		User                string `short:"u" long:"user" description:"database user"`
		Password            string `long:"password" description:"database password"`
		InteractivePassword bool   `short:"p" description:"input password interactively"`
		DefaultsExtraFile   string `long:"defaults-extra-file" description:"like mysql --defaults-extra-file for hidden password"`
		Host                string `short:"h" long:"host" default:"127.0.0.1" description:"database host"`
		Port                int    `short:"P" long:"port" default:"3306" description:"database port"`
		Socket              string `short:"S" long:"socket" description:"unix socket file"`
		Database            string `short:"d" long:"database" description:"database name"`
		Table               string `long:"table" description:"table name"`
		Charset             string `long:"charset" default:"utf8mb4" description:"connection charset"`
		Limit               int    `long:"limit" description:"query result lines limit"`

		// read from fille
		Query            string `short:"e" long:"query" description:"query read from file or command line"`
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

	var opt option
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
		return fmt.Errorf("")
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
		err := parseDefaultsExtraFile(opt.DefaultsExtraFile)
		if err != nil {
			return err
		}
	}

	if Cfg.User != "" {
		opt.User = Cfg.User
	}
	if Cfg.Password != "" {
		opt.Password = Cfg.Password
	}
	if Cfg.Charset != "" {
		opt.Charset = Cfg.Charset
	}
	if Cfg.Host != "" {
		opt.Host = Cfg.Host
	}
	if Cfg.Port != "" {
		opt.Port, err = strconv.Atoi(Cfg.Port)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
	if Cfg.Database != "" {
		opt.Database = Cfg.Database
	}

	if len(opt.Comma) > 1 || opt.Comma == "" {
		println("csv comma is rune, only one character")
		os.Exit(1)
	}

	if !strings.HasPrefix(strings.ToLower(opt.Charset), "utf") {
		opt.BOM = false
	}

	if opt.InteractivePassword {
		fmt.Print("Password:")
		password, err := gopass.GetPasswd()
		if err != nil {
			return err
		}
		opt.Password = strings.TrimSpace(string(password))
	}

	// read query interactive
	if opt.InteractiveQuery {
		// allow line separator, sql end with ';'
		fmt.Println("Query (end with '; + <Enter>'):")
		reader := bufio.NewReaderSize(os.Stdin, Cfg.MaxBufferSize)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			opt.Query = opt.Query + line
			line = strings.TrimSpace(line)
			if len(line) > 1 && line[len(line)-1] == ';' {
				break
			}
		}
	}

	// --query is empty and -table not empty generate full select query
	if opt.Table != "" && opt.Query == "" {
		if opt.Limit != 0 {
			Cfg.Server = opt.Server
			switch strings.ToLower(opt.Server) {
			case "oracle":
				opt.Query = fmt.Sprintf("SELECT * FROM %s WHERE ROWNUM <= %d", QuoteKey(opt.Table), opt.Limit)
			case "sqlserver":
				opt.Query = fmt.Sprintf("SELECT TOP %d * FROM %s", opt.Limit, QuoteKey(opt.Table))
			default:
				opt.Query = fmt.Sprintf("SELECT * FROM %s LIMIT %d", QuoteKey(opt.Table), opt.Limit)
			}
		} else {
			opt.Query = fmt.Sprintf("SELECT * FROM %s", QuoteKey(opt.Table))
		}
	}

	// test read from file
	if _, err := os.Stat(opt.Query); err == nil {
		buf, err := ioutil.ReadFile(opt.Query)
		if err == nil {
			opt.Query = string(buf)
		}
	}

	// use abs path
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if !filepath.IsAbs(opt.File) &&
		(opt.File != "" && opt.File != "stdout") {
		opt.File = filepath.Join(pwd, opt.File)
	}

	Cfg = Config{
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

		Query:                   opt.Query,
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

	// get table name from file prefix
	if Cfg.Table == "" {
		Cfg.Table = strings.Split(filepath.Base(Cfg.File), ".")[0]
	}

	if Cfg.Server == "sqlite" &&
		Cfg.Database == "" && Cfg.DSN == "" {
		println("sqlite should specified `--database DATA_FILE` arg")
		os.Exit(1)
	}

	return err
}
