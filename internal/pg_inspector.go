package internal

import (
	"database/sql"

	"github.com/kaddiya/dbi/pkg"
	_ "github.com/lib/pq"
)

type PGDBInspector struct {
	DBConn *sql.DB
}

func (pgDB PGDBInspector) GetTables() ([]*pkg.DbiTables, error) {
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

func (pgDB PGDBInspector) GetColumnsForTable(tableName string) ([]*pkg.DbiColumns, error) {
	var res []*pkg.DbiColumns

	q, err := pgDB.DBConn.Query("SELECT is_nullable,data_type,column_name FROM information_schema.columns where table_name=$1", tableName)

	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results

	for q.Next() {
		t := pkg.DbiColumns{}

		// scan
		err = q.Scan(&t.Nullable,
			&t.DataType,
			&t.ColumnName,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}
	return res, nil
}

func (pgDB PGDBInspector) GetConstraintsForTable(tableName string) ([]*pkg.DbiConstraints, error) {

	var res []*pkg.DbiConstraints

	q, err := pgDB.DBConn.Query("SELECT constraint_name,constraint_type from information_schema.table_constraints where table_name=$1", tableName)

	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results

	for q.Next() {
		t := pkg.DbiConstraints{}

		// scan
		err = q.Scan(&t.ConstraintName,
			&t.ConstraintType,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}
	return res, nil
}

func (pgDB PGDBInspector) GetKeyUsageForTable(tableName string) ([]*pkg.DbiKeyUsages, error) {

	var res []*pkg.DbiKeyUsages

	q, err := pgDB.DBConn.Query("select column_name,constraint_name from information_schema.key_column_usage where table_name=$1", tableName)

	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results

	for q.Next() {
		t := pkg.DbiKeyUsages{}

		// scan
		err = q.Scan(&t.ColumnName,
			&t.ConstraintName,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}
	return res, nil
}
