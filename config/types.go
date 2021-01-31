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

// Storage the storage configuration
type Storage struct {
	Engine    string `json:"engine"` // osfs/cos/mem ...
	Root      string `json:"root,omitempty"`
	Namespace string `json:"namespace,omitempty"`  // mem only
	SecretID  string `json:"secret_id,omitempty"`  // cos only
	SecretKey string `json:"secret_key,omitempty"` // cos only
	Bucket    string `json:"bucket,omitempty"`     // cos only (bucket name)
	Readonly  bool   `json:"readonly,omitempty"`
}

// Config Yao framework configuration
type Config struct {
	Default    map[string]string     `json:"default"`
	Database   map[string][]Database `json:"database"`
	Storage    map[string]Storage    `json:"storage"`
	Connection Connection            `json:"connection"`
}
