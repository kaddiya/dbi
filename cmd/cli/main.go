package main

import (
	"flag"
	"fmt"

	"github.com/kaddiya/dbi/internal"
)

func main() {

	protocol := flag.String("protocol", "pgsql", "protocol for the DB")
	userName := flag.String("u", "local", "user name")
	password := flag.String("p", "local", "password")
	host := flag.String("h", "localhost", "host url")
	dbName := flag.String("d", "todo", "database name")
	sslMode := flag.String("s", "disable", "sslMode")

	flag.Parse()

	dbCfgBldr := &internal.DBInspectorBuilderImpl{}

	inspector, cfgErr := dbCfgBldr.GetDbInspectorInstance(*protocol, *userName, *password, *host, *dbName, *sslMode)
	if cfgErr != nil {
		panic(cfgErr)
	}
	tbls, tblErr := inspector.ListTables()
	if tblErr != nil {
		panic(tblErr)
	}
	for _, val := range tbls {
		fmt.Println(val.Name)
	}
}
