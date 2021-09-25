# Import Data

## Table Name

If `--table` flag was specified, d18n will load data into this table.

If `--table` flag was not specified, d18n will use the file name prefix as table name. For example 'actor.csv', table name is 'actor'.

## Table definition

d18n can load schema information from runtime database or from `--schema` config file.

d18n use `SELECT * FROM ${TABLE} LIMIT 0` query check table definition from runtime database and load data into database directly. If you want to check SQL first, please don't specify database access information. Use `--schema` config to load table definition, and d18n will print SQL into stdout. You can redirect stdout to check SQL syntax or data correctness.

If import data file columns sequence mapping with table definition, edit new order in `--schema` config file can help restruct columns' sequence.

If data file only contain part of columns, not full columns, use `--schema` and `--complete-insert` may give help, other columns will use default value as table defined.

## Skip Lines

Sometimes files contain some dirty data, like table header. Use `--skip-lines` can jump first n lines from the file, load data after that.

## Foreign Key Check

MySQL supports temporary disable foreign key checks by session-level. If you what to load a table with a foreign key, you can use `--disable-foreign-key-checks` flag to disable foreign key checks. PostgreSQL, SQLite also supports this flag.

There is a more compatible way to disable foreign key checks, but more cross-database compatible. That's changing table definition, remove foreign key constraints, and add them back after loaded.
