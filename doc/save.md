# Save File

## Support File Type

Flag `--file` supports the following file types. Without the flag, d18n will print query result as ASCII table into `stdout`. If `--file stdout`, d18n will print query result as txt format into `stdout`.

* xlsx: Microsoft Office Excel. [Example](../test/actor.xlsx)
* csv: Comma Separated Values. [Example](../test/actor.csv)
* sql: SQL (Structured Query Language) file. [Example](../test/actor.sql)
* txt: plain text, separated by space. [Example](../test/actor.txt)
* tsv: Tab Separated Values. [Example](../test/actor.tsv)
* psv: Pipe Separated Values. [Example](../test/actor.psv)
* json: JSON (JavaScript Object Notation) file. [Example](../test/actor.json)
* html: HTML (HyperText Markup Language) file. [Example](../test/actor.html)

Query result print as ASCII table.

```bash
~ $ d18n --defaults-extra-file test/my.cnf --query "show databases" --verbose
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
Get rows: 7 Query cost: 3.943ms Save cost: 247µs Total Cost: 4.191ms
```

Except `xlsx`， other file formats also support print to `stdout`

```bash
~ $ d18n --defaults-extra-file test/my.cnf --query "show databases" --file csv
Database
information_schema
mysql
performance_schema
sakila
sys
test
world_x
```

Combine with [s3cmd](https://github.com/s3tools/s3cmd), [mc](https://docs.min.io/docs/minio-client-quickstart-guide.html), [rclone](https://rclone.org/), you can save select result into cloud storage ike `s3`, `hdfs`, `gcs`...

```bash
# Example for saving file into minio
~ $ d18n --defaults-extra-file test/my.cnf --query "show databases" --file csv | \
mc pipe s3/mybucket/d18n_query_result.csv
```

Save query result into a file.

```bash
~ $ d18n --defaults-extra-file test/my.cnf --query "show databases" --file result.csv
```


## Query

d18n can read the query from the file after `--query` flag. If the file does not exist, d18n reads the flag value as query SQL.

```bash
~ $ d18n --query query.sql

~ $ d18n --query "select 1"
```

## Interactive Query

d18n can be used as a simple interactive client, with `-q` flag.

```bash
~ $ d18n --defaults-extra-file test/my.cnf -q
Query (end with '; + <Enter>'):
select 1;
+---+
| 1 |
+---+
| 1 |
+---+
Query (end with '; + <Enter>'):
show databases;
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
Query (end with '; + <Enter>'):
<Ctrl>+D
EOF
```

## File Size Limitation

Default excel max size is 10MB, other file type's size with no limit. Excel size limit can be changed with `--excel-max-file-size` flag. Increasing this size will cause more memory usage.

Excel has many more limitations. For example:

```txt
MaxRows      = 1048576 // 1M
MaxColumns   = 16384   // 16K
MaxCellChars = 32767   // 32K
```

If CSV file too large, open with Excel will overflow. At this time, please open CSV file with plain text viewer like `less`, `more`.

![excel_overflow](./images/excel_overflow.png)

## Extra Statistic Information

With `--verbose` flag, d18n will print extra statistic information into `stderr`.

Extra statistic information contains rows and time costs.

```bash
~ $ d18n --defaults-extra-file test/my.cnf --query "show databases" --file result.csv --verbose
Get rows: 7 Query cost: 23.382ms Save cost: 297µs Total Cost: 23.679ms
```

## Limit lines

There are two methods to limit return lines.

1. Use `LIMIT` clause in SQL.
2. Use `--limit` flag.

```bash
~ $ d18n --defaults-extra-file test/my.cnf --database sakila --query "select * from actor limit 10"

~ $ d18n --defaults-extra-file test/my.cnf --database sakila --query "select * from actor" --limit 10
```

## BOM (Byte Of Mark)

CSV file open by Microsoft Office doesn't use UTF8 encoding by default. It will choose OS default encoding like ANSI. Specify `--bom` flag will write utf8 BOM at the head of the file, which will tell Microsoft Office to open files using UTF8 encoding.

## No Header

If you only want to save data without column names, flag `--no-header` will do this.

## NULL string

You can change `NULL` value by `--null-string` flag. e.g., nil, NULL, None.

## Check Empty

We use d18n dump data should always get data, if it returns empty set, use `--check-empty` flag, d18n will raise an error, exit code none zero. It's useful for the fail check.
