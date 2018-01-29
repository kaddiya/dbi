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

	/*
		dbCfgBldr := &internal.DBInspectorBuilderImpl{}

		inspector, cfgErr := dbCfgBldr.GetDbInspectorInstance(*protocol, *userName, *password, *host, *dbName, *sslMode)
		if cfgErr != nil {
			panic(cfgErr)
		}
		tbls, tblErr := inspector.GetTables()
		if tblErr != nil {
			panic(tblErr)
		}

		for _, val := range tbls {
			var reflectionErr error
			var colList []*pkg.DbiColumns
			var constraintsList []*pkg.DbiConstraints
			var keyUsages []*pkg.DbiKeyUsages

			colList, reflectionErr = inspector.GetColumnsForTable(val.TableName)
			constraintsList, reflectionErr = inspector.GetConstraintsForTable(val.TableName)
			keyUsages, reflectionErr = inspector.GetKeyUsageForTable(val.TableName)
			if reflectionErr != nil {
				fmt.Println("Could not get the column data for " + val.TableName + " due to " + reflectionErr.Error())
			} else {
				fmt.Println("*********************")
				fmt.Println(val.TableName)
				for _, cols := range colList {
					fmt.Println(cols.ColumnName, cols.DataType)
				}
				for _, constraints := range constraintsList {
					fmt.Println(constraints.ConstraintName, constraints.ConstraintType)
				}
				for _, keyUsages := range keyUsages {
					fmt.Println(keyUsages.ColumnName, keyUsages.ConstraintName)
				}
			}

		}*/
}
