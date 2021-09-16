# Testing

## Sakila Test Database

d18n test with [sakila](https://dev.mysql.com/doc/sakila/en/) sample database. It supports many types of databases, besides MySQL other databases can be found in https://github.com/jOOQ/sakila.

## Run Database Instance with Docker

d18n use docker test different RDBMS instances.

```bash
# community mysql latest
make docker-mysql

# community mysql 5.7
MYSQL_VERSION=5.7 make docker-mysql

# percona-server
MYSQL_RELEASE=percona/percona-server MYSQL_VERSION=latest make docker-mysql

# oracle 11g
make docker-oracle

# postgres latest
make docker-postgres

# mssql 2017-latest
make docker-mssql

# clickhouse latest
make docker-clickhouse
```
