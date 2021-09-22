# Development

After downloading all the go modules, d18n can be developed offline without the internet.

## Dependencies

d18n use `embed` module, Go version must 1.16+.

## Makefile

Makefile help d18n build & test automatically.

```bash
# build binary for current platform
make build

# build release able binary for Windows/Mac/Linux platform
make release
```

## Run database in docker

d18n test different database with docker instances. All the Makefile are located in test directory and included by outside Makefile.

```bash
# stop all running container
make docker-stop

# run MySQL instance
make docker-mysql

# connect to MySQL interactively
make docker-connect

# run Oracle instance
make docker-oracle

# connect to Oracle interactively
make docker-sqlplus

# run SQL Server instance
make docker-mssql

# connect to SQL Server interactively
make docker-sqlcmd

# run PostgreSQL instance
make docker-postgres

# connect to PostgreSQL interactively
make docker-psql

# run ClickHouse instance
make docker-clickhouse

# connection to ClickHouse interactively
make docker-ck-client
```

Note: d18n auto use [podman](https://podman.io/) first if installed.

## Test

```bash
# run all MySQL test cases before release
make ci

# run all unit test cases
make test

# run all test cases and output coverage information
make cover

# run database specified cases
make test-mysql
make test-oracle
make test-mssql
make test-postgres
make test-clickhouse

```

## Others

```bash
# code format
make fmt

# go mod tidy
make tidy
```
