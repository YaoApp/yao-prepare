package database

import (
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
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
	policy   dbresolver.RandomPolicy
	db       *gorm.DB
}
