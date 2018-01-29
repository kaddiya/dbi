package pkg

import "database/sql"

type DBInspector interface {
	GetTables() ([]*DbiTables, error)
	GetColumnsForTable(string) ([]*DbiColumns, error)
}

type DBConfig struct {
	DBName   string
	SSLMode  bool
	UserName string
	Password string
	Host     string
	Protocol string
}

type DbiTables struct {
	TableCatalog              string
	TableSchema               string
	TableName                 string
	TableType                 string
	SelfReferencingColumnName sql.NullString
	ReferenceGeneration       sql.NullString
	UserDefinedTypeCatalog    sql.NullString
	UserDefinedTypeSchema     sql.NullString
	UserDefinedTypeName       sql.NullString
	IsInsertableInto          string
	IsTyped                   string
	CommitAction              sql.NullString
}

type DbiColumns struct {
	Nullable      string
	DataType      string
	ColumnDefault string
	ColumnName    string
}
