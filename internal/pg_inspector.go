package internal

import (
	"database/sql"

	"github.com/kaddiya/dbi/pkg"
)

type PGDBInspector struct {
	DBConn *sql.DB
}

func (pgDB PGDBInspector) ListTables() ([]*pkg.DbiTables, error) {
	var res []*pkg.DbiTables
	q, err := pgDB.DBConn.Query("SELECT * FROM information_schema.tables  where table_schema='public'")

	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results

	for q.Next() {
		t := pkg.DbiTables{}

		// scan
		err = q.Scan(&t.TableCatalog,
			&t.TableSchema,
			&t.TableName,
			&t.TableType,
			&t.SelfReferencingColumnName,
			&t.ReferenceGeneration,
			&t.UserDefinedTypeCatalog,
			&t.UserDefinedTypeSchema,
			&t.UserDefinedTypeName,
			&t.IsInsertableInto,
			&t.IsTyped,
			&t.CommitAction)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}
