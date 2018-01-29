package internal

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PGDBInspector struct {
	DBConn *sql.DB
}

func (pgDB PGDBInspector) GetTables() ([]*DbiTables, error) {
	var res []*DbiTables
	q, err := pgDB.DBConn.Query("SELECT * FROM information_schema.tables  where table_schema='public'")

	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results

	for q.Next() {
		t := DbiTables{}

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

func (pgDB PGDBInspector) GetColumnsForTable(tableName string) ([]*DbiColumns, error) {
	var res []*DbiColumns

	q, err := pgDB.DBConn.Query("SELECT is_nullable,data_type,column_name FROM information_schema.columns where table_name=$1", tableName)

	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results

	for q.Next() {
		t := DbiColumns{}

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

func (pgDB PGDBInspector) GetConstraintsForTable(tableName string) ([]*DbiConstraints, error) {

	var res []*DbiConstraints

	q, err := pgDB.DBConn.Query("SELECT constraint_name,constraint_type from information_schema.table_constraints where table_name=$1", tableName)

	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results

	for q.Next() {
		t := DbiConstraints{}

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

func (pgDB PGDBInspector) GetKeyUsageForTable(tableName string) ([]*DbiKeyUsages, error) {

	var res []*DbiKeyUsages

	q, err := pgDB.DBConn.Query("select column_name,constraint_name from information_schema.key_column_usage where table_name=$1", tableName)

	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results

	for q.Next() {
		t := DbiKeyUsages{}

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
