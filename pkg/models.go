package pkg

type DBInspector interface {
	ListTables() ([]*DbiTables, error)
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
	Name string
}
