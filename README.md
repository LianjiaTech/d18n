# Introduction

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)

![logo](./logo_64x64.png)

`d18n` is a [numeronym](https://en.wikipedia.org/wiki/Numeronym) short for "data-desensitization", sounds like "d-eighteen-n".

As its name says, d18n can mask data to make it desensitized. In addition, d18n can do many other things.

* d18n is a portable RDBMS cmd client. e.g., MySQL, PostgreSQL, Oracle, SQL Server ...
* save query result into a file, e.g., `xlsx`, `csv`, `txt`, `sql`, `html`, `json` ...
* detect sensitive info (like PII) from a file or a SQL query.
* import data from files into different types of databases.
* lint data file, to check if its format is compatible before import it into some database.

It can be used as a portable cmd client or imported as a package by other tools.

For more details and latest updates, see [doc](./doc/toc.md) and [release](https://github.com/LianjiaTech/d18n/releases) notes.

## Build

d18n develop with [Golang](https://golang.org/) 1.16+, please install first.

```bash
git clone github.com/LianjiaTech/d18n
cd d18n

# Mac or Linux
make build

# Windows
go build -o d18n cmd\d18n\d18n.go
```

## Cross-platform compile

Golang support many

```bash
~ $ make release
...

~ $ ls release
d18n.darwin-amd64  d18n.darwin-arm64  d18n.linux-amd64  d18n.windows-amd64
```

## [Quick Start](./doc/quickstart.md)

Simple Example

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

## License

d18n is under the Apache 2.0 license. See the [LICENSE](./LICENSE) file for details.
