package config

// Database the database configuration
type Database struct {
	Name     string `json:"name"`
	DSN      string `json:"dsn"`
	Driver   string `json:"driver"` // MySQL,PostgreSQL,SQL Server,SQLite,Clickhouse
	Readonly bool   `json:"readonly,omitempty"`
}

// Connection the connection pool configuration
type Connection struct {
	ConnMaxIdleTime uint32 `json:"max_idle_time,omitempty"`
	ConnMaxLifetime uint32 `json:"max_life_time,omitempty"`
	MaxIdleConns    uint32 `json:"max_idle_conn,omitempty"`
	MaxOpenConns    uint32 `json:"max_open_conn,omitempty"`
}

// Config Yao framework config
type Config struct {
	Default    map[string]string     `json:"default"`
	Database   map[string][]Database `json:"database"`
	Connection Connection            `json:"connection"`
}
