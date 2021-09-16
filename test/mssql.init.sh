#!/bin/bash

# Wait to be sure that SQL Server came up
for i in {1..50};
do
    /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P 'yourStrong(!)Password' -i /docker-entrypoint-initdb.d/mssql.init.sql >/dev/null 2>&1
    if [ $? -eq 0 ]
    then
	echo ""
        echo "mssql.init.sql completed"
        break
    else
        echo -n "."
        sleep 1
    fi
done
