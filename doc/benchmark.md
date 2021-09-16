# Benchmark

Notice: This benchmark data were results run on lap-top computer, not online machine. Find an old machine run test cases may be more helpful for optimization. The result show the slowest time, and easy for time compare.

## Generate test data

```bash
sudo apt install sysbench

sysbench --mysql-host=127.0.0.1 \
         --mysql-port=3306 \
         --mysql-user=root \
         --mysql-db=test \
         /usr/share/sysbench/oltp_common.lua \
         --tables=1 \
         --table_size=1000000 \
         prepare
```

## Benchmark backup file

Use `mysqldump` backup table.

```bash
~ $ time mysqldump -h 127.0.0.1 -u root --databases test --tables sbtest1 -r sbtest1.sql

real    0m2.880s
user    0m2.079s
sys     0m0.413s
```

Use `d18n` backup table.

```bash
~ $ time d18n --defaults-extra-file test/my.cnf --database test --table sbtest1 --file sbtest2.sql --verbose
Get rows: 999900 Query cost: 1.550356ms Save cost: 5.509046163s Total Cost: 5.510596605s

real    0m5.750s
user    0m5.532s
sys     0m0.243s
```

Use `--extended-insert` flag with `d18n`

```bash
~ $ time d18n --defaults-extra-file test/my.cnf --database test --table sbtest1 --file sbtest3.sql --verbose --extended-insert 100
Get rows: 1000000 Query cost: 2.049568ms Save cost: 7.008905042s Total Cost: 7.010954668s

Get rows: 999900 Query cost: 1.553367ms Save cost: 5.517253703s Total Cost: 5.518807155s

real    0m5.746s
user    0m5.448s
sys     0m0.317s
```

## Benchmark import file

Use `mysql` client import SQL file.

```bash
~ $ time mysql -h 127.0.0.1 -u root --database test -f < ./sbtest1.sql

real    0m36.333s
user    0m2.494s
sys     0m0.261s
```

Use `d18n` import SQL file.

```bash
~ $ time d18n --defaults-extra-file test/my.cnf --database test --table sbtest1 --file sbtest2.sql --import

real    25m4.975s
user    1m24.025s
sys     1m59.032s
```

Use `mysql` client import d18n SQL file.

```bash
~ $ time mysql -h 127.0.0.1 -u root --database test -f < ./sbtest2.sql
real    21m35.807s
user    0m39.472s
sys     0m43.766s
```

Use `d18n` client import extended-insert SQL file.

```bash
time d18n --defaults-extra-file test/my.cnf --database test --table sbtest1 --file sbtest3.sql --import --verbose
Import Rows: 10000 Total Cost: 35.764419035s

real    0m36.063s
user    0m2.568s
sys     0m1.110s
```

Use `mysql` client import d18n extended-insert SQL file.

```bash
~ $ time mysql -h 127.0.0.1 -u root --database test -f < ./sbtest3.sql
real    0m37.039s
user    0m2.153s
sys     0m0.558s
```

## SELECT INTO OUTFILE

Backup table into csv file with MySQL native syntax.

```bash
mysql >
SELECT *
  INTO OUTFILE 'sbtest1.csv'
  FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"'
  LINES TERMINATED BY '\n'
  FROM sbtest1;

Query OK, 1000000 rows affected (1.62 sec)
```

Backup table into csv use `d18n`

```bash
$ time d18n --defaults-extra-file test/my.cnf --database test --table sbtest1 --file sbtest.csv --verbose
Get rows: 1000000 Query cost: 1.298027ms Save cost: 5.732692558s Total Cost: 5.733990955s

real  0m5.978s
user  0m4.454s
sys   0m0.841s
```

## LOAD DATA INFILE

Load csv file with MySQL native syntax.

```bash
mysql >
LOAD DATA INFILE 'sbtest1.csv'
INTO TABLE sbtest1
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n';

Query OK, 1000000 rows affected (15.52 sec)
Records: 1000000  Deleted: 0  Skipped: 0  Warnings: 0
```

Load csv file use d18n

```bash
$ time d18n --defaults-extra-file test/my.cnf --database test --table sbtest1 --file sbtest.csv --extended-insert 100 --import

real  0m55.022s
user  0m20.084s
sys   0m1.733s
```
