# Database

## System Support

d18n develop by Go language, so it's cross-platform, you deploy d18n in any OS
that Golang support.

d18n doesn't use `CGO` database driver, so it's portable, without any dynamic libraries
dependence.

## Database Support

d18n wants to support every database use SQL. Please report issue in github, if you find a new database or a bug.

* MySQL: Community, MariaDB, TiDB, Percona
* PostgreSQL: Community, CockroachDB
* Oracle
* SQLite3
* SQL Server
* ClickHouse
* CSV: use [csvq](github.com/mithrandie/csvq-driver) library for csv file query

## Use d18n as a simple client

d18n can be used as a simple query client. It supports many types of databases.

```bash
~ $ d18n --defaults-extra-file test/my.cnf --server mysql --query "show databases"
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

~ $ d18n --server postgres --database postgres --query 'SELECT datname FROM pg_database WHERE datistemplate = false;' --user postgres -p --port 5432
Password: *******
+----------+
| DATNAME  |
+----------+
| postgres |
+----------+
```

## DSN

There are three methods for connecting databases.

First: d18n support `--defaults-extra-file` flag like mysql client, and `--defaults-extra-file` must be the first flag.

my.cnf example, only support the following keyword.

```ini
[client]
user=root
password=******
default-character-set=utf8mb4
```

Second: `--host`, `--user`, `--password`, `--port`, `--database`, `-p`, `--socket`, `--charset`, `--ansi-quotes`

`-p` flag for the password, d18n will read password interactively.

```bash
~ $ d18n --host 127.0.0.1 --user root --password "******" --port 3306 --query "show databases"
```

Third: use `--dsn` flag to overwrite other database flags.

```bash
~ $ d18n --dsn "user:password@tcp(127.0.0.1:3306)/sakila" --query "show databases"
```

### MySQL

d18n connect mysql use `github.com/go-sql-driver/mysql` driver, [full DSN reference](https://github.com/go-sql-driver/mysql#dsn-data-source-name).

### PostgreSQL

d18n connect PostgreSQL use `github.com/lib/pq` driver, [full DSN reference](https://github.com/lib/pq/blob/master/doc.go).

### Oracle

d18n connect oracle use `github.com/sijms/go-ora/v2` driver, [full DSN reference](https://github.com/sijms/go-ora#servers-url-options)

### SQLite3

d18n connect SQLite3 use `modernc.org/sqlite` driver, [full DSN reference](https://pkg.go.dev/modernc.org/sqlite?utm_source=godoc)

### SQL Server

d18n connect SQL Server use `github.com/denisenkom/go-mssqldb` driver, [full DSN reference](https://github.com/denisenkom/go-mssqldb#connection-parameters-and-dsn)

### ClickHouse

d18n connect ClickHouse use `github.com/ClickHouse/clickhouse-go` driver, [full DSN reference](https://github.com/ClickHouse/clickhouse-go#dsn)

### CSV

d18n connect CSV use `github.com/mithrandie/csvq-driver` driver, [full DSN reference](https://github.com/mithrandie/csvq-driver#data-source-name)

## INSERT/REPLACE

`INSERT` syntax is mostly cross-database compatible.

`INSERT` syntax:

```sql
INSERT INTO table_name
VALUES (value1, value2, value3, ...);
```

`--complete-insert` flag will auto-complete column name.

```sql
INSERT INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);
```

```bash
~ $ d18n --defaults-extra-file test/my.cnf --server mysql --query "select * from actor" --database sakila --complete-insert --file actor.sql
```

d18n also supports saving SQL use `REPLACE` syntax. Choose `--replace` flag to change SQL syntax. Notice: `REPLACE` are not cross database type compatible.

`REPLACE` syntax:

```sql
REPLACE INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);
```

```bash
~ $ d18n --defaults-extra-file test/my.cnf --server mysql --query "select * from actor" --database sakila --replace --file actor.replace.sql
```

Notice: d18n support `--extended-insert` flag like mysqldump. d18n will save SQL one by one by default, without values merge. Use `--extended-insert` flag can merge multi values into one SQL, which will decrease file size, and make import speed faster.

If `--extended-insert` too large will cause bufio overflow. Please consider increasing `--max-buffer-size`.

## UPDATE

d18n also supports saving the result as `UPDATE` SQL. It's useful for column backup.

`UPDATE` syntax:

```sql
UPDATE table_name
SET column1 = value1, column2 = value2, ...
WHERE condition;
```

Primary keys after `--update` flag should join with a comma(`,`). Column names are case insensitive. d18n supports multi-column primary key.

```bash
~ $ d18n --defaults-extra-file test/my.cnf --server mysql --query "select * from actor" --database sakila --update actor_id --file actor.update.sql
```

## Load Schema

When d18n import plan data into database, it need columns name and columns data type. So give a schema file can help d18n determine column order and which column is which data type. If no `--schema` flag specified it will use table original columns order import data.

`--schema` flag will load the schema config file. This config file contains only one table schema info, which can be `CREATE TABLE` SQL or plain text config. Schema support MySQL, Oracle, PostgreSQL, SQL Server ...

e.g., --schema schema.txt

```text
# comment line
actor_id SMALLINT
first_name VARCHAR
last_name VARCHAR
last_update TIMESTAMP
```

e.g., --schema mysql.schema.sql

```sql
# SHOW CREATE TABLE `actor`
CREATE TABLE actor (
  actor_id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  first_name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NOT NULL,
  last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (actor_id),
  KEY idx_actor_last_name (last_name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

e.g., --schema oracle.schema.sql

```sql
CREATE TABLE actor (
  actor_id numeric NOT NULL ,
  first_name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NOT NULL,
  last_update DATE NOT NULL,
  CONSTRAINT pk_actor PRIMARY KEY (actor_id)
);
```

## HEX BLOB

`--hex-blob`

Dump binary columns using hexadecimal notation (for example, 'abc' becomes 0x616263). The affected data types are `BLOB`, `BINARY`, `VARBINARY` types, BIT, all spatial data types, and other non-binary data types when used with the binary character set.

Spatial data types suggest use `--hex-blob` flag, e.g., `GEOMETRY`, `POINT`, `LINESTRING`, `POLYGON`, `MULTIPOINT`, `MULTILINESTRING`, `MULTIPOLYGON`, `GEOMETRYCOLLECTION`.
