# Use d18n as a comment client

d18n is portable and compatible with many databases. Why not use d18n as a comment client?

Here are some examples.

```bash
# `-p` input password interactively, `-q` input query interactively
./bin/d18n --host 127.0.0.1 --user username -p -q
Password:
mysql > select 1;
+---+
| 1 |
+---+
| 1 |
+---+
mysql > <Ctrl+D>
EOF

# use login-path encrypt login info
./bin/d18n --login-path=username -q
mysql > show databases;
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
mysql > <Ctrl+D>
EOF

# Vertical output
show databases\G
+---------------------------+---------------------------+
| ********* Row 1 ********* | ********* Row 1 ********* |
| Database                  | information_schema        |
| ********* Row 2 ********* | ********* Row 2 ********* |
| Database                  | mysql                     |
| ********* Row 3 ********* | ********* Row 3 ********* |
| Database                  | performance_schema        |
| ********* Row 4 ********* | ********* Row 4 ********* |
| Database                  | sakila                    |
| ********* Row 5 ********* | ********* Row 5 ********* |
| Database                  | sys                       |
| ********* Row 6 ********* | ********* Row 6 ********* |
| Database                  | test                      |
| ********* Row 7 ********* | ********* Row 7 ********* |
| Database                  | world_x                   |
+---------------------------+---------------------------+

# login Oracle
# d18n support mysql, oracle, mssql, postgres, clickhouse ...
./bin/d18n --server oracle --host 127.0.0.1 --port 1521 --user user_name --database service_name -q -p
Password:
oracle > select 1 from dual;
+---+
| 1 |
+---+
| 1 |
+---+
oracle >

# wrap d18n with `rlwrap` for readline
# it can remember input history and many other keyboard shortcuts.
rlwrap ./bin/d18n --login-path=username -q


# customize prompt
--prompt "\x1b[31m master \x1b[0m> "
```
