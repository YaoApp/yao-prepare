package config

// Database the database configuration
type Database struct {
	Name     string `json:"name"`
	DSN      string `json:"dsn"`
	Driver   string `json:"driver"` // MySQL,PostgreSQL,SQL Server,SQLite,Clickhouse
	Readonly bool   `json:"readonly,omitempty"`
}

// Config Yao framework config
type Config struct {
	Default  map[string]string `json:"default"`
	Database []Database        `json:"database"`
}
