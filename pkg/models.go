package pkg

type DBInspectorAPI interface {
	GetDatabaseMetadata(*DBConfig) ([]*Table, error)
}

type DBInspectorImpl struct {
}

type DBConfig struct {
	DBName   string
	SSLMode  string
	UserName string
	Password string
	Host     string
	Protocol string
}

type Column struct {
	ColumnName     string
	ConstraintType string
	ConstraintName string
}

type Table struct {
	TableName string
	Columns   []*Column
}
