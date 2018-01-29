package main

import (
	"flag"
	"fmt"

	"github.com/kaddiya/dbi/pkg"
)

func main() {

	protocol := flag.String("protocol", "pgsql", "protocol for the DB")
	userName := flag.String("u", "local", "user name")
	password := flag.String("p", "local", "password")
	host := flag.String("h", "localhost", "host url")
	dbName := flag.String("d", "local", "database name")
	sslMode := flag.String("s", "disable", "sslMode")

	flag.Parse()

	inspector := &pkg.DBInspectorImpl{}
	data, r := inspector.GetDatabaseMetadata(&pkg.DBConfig{
		DBName:   *dbName,
		SSLMode:  *sslMode,
		UserName: *userName,
		Password: *password,
		Host:     *host,
		Protocol: *protocol,
	})

	if r != nil {
		fmt.Println(r.Error())
	}
	for _, tbls := range data {
		fmt.Println("***************")
		fmt.Println(tbls.TableName)
		fmt.Println(len(tbls.Columns))
		for _, cols := range tbls.Columns {
			fmt.Println(cols.ColumnName, cols.ConstraintName, cols.ConstraintType)
		}
		fmt.Println("-----------------")
	}
}
