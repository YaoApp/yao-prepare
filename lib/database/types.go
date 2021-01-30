package database

import (
	"github.com/yaoapp/yao/config"
	"gorm.io/gorm"
)

type conn struct {
	dialector gorm.Dialector
	name      string
}

// Pool the connections pool
type Pool struct {
	main     gorm.Dialector
	sources  []conn
	replicas []conn
	option   config.Connection
	db       *gorm.DB
}
