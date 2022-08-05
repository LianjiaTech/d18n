module github.com/LianjiaTech/d18n

go 1.16

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.4.0
	github.com/ClickHouse/clickhouse-go v1.4.5
	github.com/andrewarchi/gocipher v0.0.0-20201117064119-d399f25a1970
	github.com/antlr/antlr4 v0.0.0-20181218183524-be58ebffde8e // indirect
	github.com/auxten/postgresql-parser v1.0.0 // indirect
	github.com/brianvoe/gofakeit/v6 v6.5.0
	github.com/bykof/gostradamus v1.0.4
	github.com/capitalone/fpe v1.2.1
	github.com/denisenkom/go-mssqldb v0.10.0
	github.com/dnnrly/abbreviate v1.5.2
	github.com/dolmen-go/mylogin v0.0.0-20211007211255-18ac7793a122
	github.com/go-ego/gse v0.67.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/differential-privacy/go v0.0.0-20210713105217-8da48001ccbd
	github.com/google/gxui v0.0.0-20151028112939-f85e0a97b3a4 // indirect
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/jessevdk/go-flags v1.5.0
	github.com/jmrobles/h2go v0.5.0
	github.com/json-iterator/go v1.1.11
	github.com/kr/pretty v0.2.1
	github.com/lib/pq v1.10.2
	github.com/mithrandie/csvq-driver v1.4.3
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pingcap/parser v0.0.0-20210525032559-c37778aff307
	github.com/pingcap/tidb v1.1.0-beta.0.20210601085537-5d7c852770eb
	github.com/pingcap/tipb v0.0.0-20210601083426-79a378b6d1c4 // indirect
	github.com/prestodb/presto-go-client v0.0.0-20201204133205-8958eb37e584
	github.com/sijms/go-ora/v2 v2.4.28
	github.com/taozle/go-hive-driver v0.0.0-20181206100408-79951111cb07
	github.com/tealeg/xlsx/v3 v3.2.3
	github.com/tjfoc/gmsm v1.3.2
	github.com/wumansgy/goEncrypt v0.0.0-20201114063050-efa0a0601707
	github.com/zach-klippenstein/goregen v0.0.0-20160303162051-795b5e3961ea
	golang.org/x/exp v0.0.0-20210220032938-85be41e4509f // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/yaml.v2 v2.4.0
	modernc.org/sqlite v1.11.2
)

// fix potential security issue(CVE-2020-26160) introduced by indirect dependency.
replace github.com/dgrijalva/jwt-go => github.com/form3tech-oss/jwt-go v3.2.6-0.20210809144907-32ab6a8243d7+incompatible
