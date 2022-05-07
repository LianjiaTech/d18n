# File Lint

d18n supports a new feature about data file lint. Check data format use `--lint` flag can give more assurance.

## Lint Rules

| Rule Name           | Description                                                            | Level |
| :------------------ | :--------------------------------------------------------------------- | :---- |
| RaggedRows          | Rows in the file don't have the same number of columns.              | ERROR |
| UnMatchHeader       | Header number not match value number.                                  | ERROR |
| UndeclaredHeader    | First line in file can't be used as column names.                      | ERROR |
| DuplicateColumnName | Column names aren't unique.                                            | ERROR |
| CheckOptions        | Cells less or equal than 1 .                                           | WARN  |
| CellSpace           | Cell leading or ending with space.                                     | WARN  |
| LineBreaks          | Line breaks are not the same as define.                                | ERROR |
| UnclosedQuote       | There are any unclosed quotes in line.                                 | ERROR |
| LeadingSpace        | Line leading with space.                                               | WARN  |
| BlankRows           | There are any blank rows.                                              | WARN  |
| Whitespace          | There is any whitespace between commas and double quotes around cells. | WARN  |
| CommentRows         | There are any comment rows.                                            | WARN  |

## Use

```shell
$ d18n --lint --file test/actor.csv
ok

$ d18n --lint --file test/TestCSVLint.wrong.csv
Line: 1, Column: 14, ERROR: First line in file can't be used as column names.

$ d18n --lint --file test/TestXLSXLint.wrong.xlsx
Line: 1, Column: 2, ERROR: First line in file can't be used as column names.
Line: 1, Column: 7, ERROR: Column names aren't unique.

$ d18n --lint --file test/TestJSONLint.wrong.json
Line: 1, Column: 4, ERROR: Column names aren't unique.
Line: 1, Column: 1, ERROR: First line in file can't be used as column names.

```

You can also use the `--verbose` flag to see more information

```shell
$ d18n --lint --file test/TestCSVLint.wrong.csv --verbose

$ d18n --lint --file test/TestCSVLint.wrong.csv --verbose
Line: 1, Column: 14, ERROR: First line in file can't be used as column names.

File Size: 469
Row Count(Include Header): 3 Cell Count: 15 Error Count: 1 Time Cost: 172µs

$ d18n --lint --file test/TestXLSXLint.wrong.xlsx --verbose
Line: 1, Column: 7, ERROR: Column names aren't unique.
Line: 1, Column: 1, ERROR: First line in file can't be used as column names.

File Size: 5888
Row Count(Include Header): 2 Cell Count: 11 Error Count: 2 Time Cost: 3.134ms

$ d18n --lint --file test/TestJSONLint.wrong.json --verbose
Line: 1, Column: 2, ERROR: First line in file can't be used as column names.
Line: 1, Column: 4, ERROR: Column names aren't unique.

File Size: 159
Row Count(Include Header): 2 Cell Count: 8 Error Count: 2 Time Cost: 461µs

```

You can easily use d18n to lint your files. It supports csv, tsk, psi, txt, json, sql, html and xlsx file formats. It distinguishes them by file suffixes. It supports the following file types.

* xlsx: Microsoft Office Excel. [Example](../test/actor.xlsx)
* csv: Comma Separated Values. [Example](../test/actor.csv)
* sql: SQL (Structured Query Language) file. [Example](../test/actor.sql)
* txt: plain text, separated by space. [Example](../test/actor.txt)
* tsv: Tab Separated Values. [Example](../test/actor.tsv)
* psv: Pipe Separated Values. [Example](../test/actor.psv)
* json: JSON (JavaScript Object Notation) file. [Example](../test/actor.json)
* html: HTML (HyperText Markup Language) file. [Example](../test/actor.html)

We are also trying to make it support more formats

## More flag

`--max-buffer-size`: If your row data is too large, it may cause bufio overflow. Please consider adding `--max-buffer-size`,defalut: 65535

`--no-header`: You can use it to skip the header check

`--comments`: You can use it to customize your comments symbols

`--server`: Different databases have different restrictions on field length. You can use this parameter to specify your database,support: mysql, postgres, sqlite, oracle, sqlserver".Default mysql.

`--ansi-quotes`: Enable ANSI_QUOTES, if you are not using a MySQL database, you can turn on the parameters
