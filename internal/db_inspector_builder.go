package internal

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/knq/dburl"
)

type DBInspectorBuilder interface {
	GetDbInspectorInstance(string, string, string, string, string, string) (*DBInspectorBuilder, error)
}

type DBInspectorBuilderImpl struct {
}

func (a *DBInspectorBuilderImpl) GetDbInspectorInstance(protocol, userName, password, host, dbName, sslMode string) (DBInspectorGranularInspector, error) {
	connString := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s", protocol, userName, password, host, dbName, sslMode)
	fmt.Println(connString)
	u, e := dburl.Open(connString)
	if e != nil {
		return nil, e
	}
	inspector := PGDBInspector{
		DBConn: u,
	}
	return inspector, nil
}
