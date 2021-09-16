# Import Data

## Table Name

If `--table` flag was specified, d18n will load data into this table.

If `--table` flag was not specified, d18n will use the file name prefix as table name. For example 'actor.csv', table name is 'actor'.

## Skip Lines

Sometimes files contain some dirty data, like table header. Use `--skip-lines` can jump first n lines from the file, load data after that.

## Foreign Key Check

MySQL supports temporary disable foreign key checks by session-level. If you what to load a table with a foreign key, you can use `--disable-foreign-key-checks` flag to disable foreign key checks. PostgreSQL, SQLite also supports this flag.

There is a more compatible way to disable foreign key checks, but more cross-database compatible. That's changing table definition, remove foreign key constraints, and add them back after loaded.
