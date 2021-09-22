# Quick Start

## As a cmd client

```bash
~ $ d18n --defaults-extra-file test/my.cnf --query "show databases"
+--------------------+
|      DATABASE      |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sakila             |
| sys                |
| test               |
| world_x            |
+--------------------+
```

my.cnf

```ini
[client]
user=root
#password=
default-character-set=utf8mb4
```

Full help please check [db](db.md) doc.

## Save query result into a file

```bash
~ $ d18n --host 127.0.0.1 --port 3306 --database sakila --user username --file actor.xlsx --query "select * from actor" -p
Password:<hidden input>
```

Notice: If the output file has existed, d18n will truncate the file first!

Full help please check [save](save.md) doc.

## Detect sensitive info

```bash
~ $ d18n --defaults-extra-file test/my.cnf --database sakila --table actor --limit 10 --detect
{
  "actor_id": null,
  "first_name": [
    "name"
  ],
  "last_name": [
    "name"
  ],
  "last_update": null
}
```

Full help please check [detect](detect.md) doc.

## Data masking

Data masking config file support `csv`, `txt` format, first argument is column name, it's case insensitive, second argument is data mask function, other arguments for function.

`mask.csv`

```csv
# column_name,mask_func,arg1,...
LAST_NAME,smokeleft,3,"x"
```

```bash
~ $ d18n --defaults-extra-file test/my.cnf \
  --mask test/mask.csv \
  --database sakila \
  --query "select * from actor where actor_id limit 10"
+----------+------------+--------------+---------------------+
| ACTOR ID | FIRST NAME |  LAST NAME   |     LAST UPDATE     |
+----------+------------+--------------+---------------------+
|        1 | PENELOPE   | xxxNESS      | 2006-02-15 04:34:33 |
|        2 | NICK       | xxxLBERG     | 2006-02-15 04:34:33 |
|        3 | ED         | xxxSE        | 2006-02-15 04:34:33 |
|        4 | JENNIFER   | xxxIS        | 2006-02-15 04:34:33 |
|        5 | JOHNNY     | xxxLOBRIGIDA | 2006-02-15 04:34:33 |
|        6 | BETTE      | xxxHOLSON    | 2006-02-15 04:34:33 |
|        7 | GRACE      | xxxTEL       | 2006-02-15 04:34:33 |
|        8 | MATTHEW    | xxxANSSON    | 2006-02-15 04:34:33 |
|        9 | JOE        | xxxNK        | 2006-02-15 04:34:33 |
|       10 | CHRISTIAN  | xxxLE        | 2006-02-15 04:34:33 |
+----------+------------+--------------+---------------------+
```

Full help please check [mask](mask.md) doc.

## Import data from a file

```bash
~ $ d18n --defaults-extra-file test/my.cnf --file test/actor.csv --replace --import --database sakila --disable-foreign-key-checks --verbose
Skip Lines: 1
Import Rows: 202 Total Cost: 259.199108ms
```

Full help please check [import](import.md) doc.

## Lint data file

```bash
~ $ d18n --file test/actor.csv --lint --verbose
ok

File Size: 7441
Row Count(Include Header): 201 Cell Count: 4 Error Count: 0 Time Cost: 887.566µs
```

Full help please check [lint](lint.md) doc.

## Preview Excel file from cmd line

```bash
~ $ d18n --file test/actor.xlsx --preview 10
actor_id	first_name	last_name	last_update
1	PENELOPE	GUINESS	2/15/06 04:34
2	NICK	WAHLBERG	2/15/06 04:34
3	ED	CHASE	2/15/06 04:34
4	JENNIFER	DAVIS	2/15/06 04:34
5	JOHNNY	LOLLOBRIGIDA	2/15/06 04:34
6	BETTE	NICHOLSON	2/15/06 04:34
7	GRACE	MOSTEL	2/15/06 04:34
8	MATTHEW	JOHANSSON	2/15/06 04:34
9	JOE	SWANK	2/15/06 04:34
```

Full help please check [preview](preview.md) doc.

## Full Usage

```text
Usage:
  d18n

Application Options:
  -v, --verbose                     verbose mode
      --help                        Show this help message
      --server=                     server type, support: mysql, postgres, sqlite, oracle, sqlserver, clickhouse (default: mysql)
      --dsn=                        formatted data source name
  -u, --user=                       database user
      --password=                   database password
  -p                                input password interactively
      --defaults-extra-file=        like mysql --defaults-extra-file for hidden password
  -h, --host=                       database host (default: 127.0.0.1)
  -P, --port=                       database port (default: 3306)
  -S, --socket=                     unix socket file
  -d, --database=                   database name
      --table=                      table name
      --charset=                    connection charset (default: utf8mb4)
      --limit=                      query result lines limit
  -e, --query=                      query read from file or command line
  -q                                input query interactively
  -f, --file=                       input/output file
      --schema=                     schema config file. support: sql, txt
      --mask=                       data masking config file. support: csv, psv, tsv format
      --cipher=                     cipher config file. support: yaml
      --sensitive=                  sensitive detection config file. support: yaml
      --print-cipher                print or auto-generate cipher
      --print-config                print config
      --preview=                    preview result file, print first N lines (default: 0)
      --lint                        lint file
      --import                      import file into database
      --detect                      detect sensitive info from data
      --watermark=                  watermark in export file. support: html, xlsx
      --check-empty                 check query result, if empty raise error
      --replace                     generate sql use replace into syntax, only support MySQL and SQLite
      --update=                     update primary key, separate by comma, case insensitive
      --complete-insert             complete insert with columns name
      --hex-blob=                   need hex encoding columns, separate by comma, case insensitive
      --ignore-columns=             import file ignore columns, separated by comma
      --extended-insert=            use multiple-row INSERT syntax that include several values list (default: 1)
      --ansi-quotes                 enable ANSI_QUOTES
      --disable-foreign-key-checks  disable foreign key checks
      --bom                         csv file with UTF8 BOM
      --excel-max-file-size=        excel max file size, limit by memory
      --lint-level=                 file lint level (default: error)
      --ignore-blank                ignore blank lines or columns when import file
      --comma=                      csv comma char (default: ,)
      --no-header                   no header line, only data lines
      --comments=                   support comment characters, multiple comment split by comma (default: #,--)
      --skip-lines=                 skip first N lines (default: 0)
      --rand-seed=                  random seed, default: current unix nano timestamp
      --max-buffer-size=            bufio MaxScanTokenSize
      --null-string=                NULL string write into file. e.g., NULL, nil, None, "" (default: NULL)
```

Note: d18n use `github.com/jessevdk/go-flags` package. It also supports Windows-style. Running in Windows, POSIX-style option will do well yet.

```text
Options with short names (/v)
Options with long names (/verbose)
Windows-style options with arguments use a colon as the delimiter
```
