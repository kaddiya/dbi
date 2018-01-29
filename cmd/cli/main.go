package main

import (
	"flag"

	"github.com/kaddiya/dbi/pkg"
)

func main() {

	protocol := flag.String("protocol", "pgsql", "protocol for the DB")
	userName := flag.String("u", "local", "user name")
	password := flag.String("p", "local", "password")
	host := flag.String("h", "localhost", "host url")
	dbName := flag.String("d", "todo", "database name")
	sslMode := flag.String("s", "disable", "sslMode")

	flag.Parse()

	inspector := &pkg.DBInspectorImpl{}
	inspector.GetDatabaseMetadata(&pkg.DBConfig{
		DBName:   *dbName,
		SSLMode:  *sslMode,
		UserName: *userName,
		Password: *password,
		Host:     *host,
		Protocol: *protocol,
	})
}
