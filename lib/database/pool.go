package database

import (
	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/arr"
	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// Add add a new source or replicate
func (pool *Pool) Add(setting config.Database) {
	switch setting.Driver {
	case "MySQL":
		pool.AddMysql(setting.Name, setting.DSN, setting.Readonly)
	}
}

// AddMysql add a MySQL DSN
func (pool *Pool) AddMysql(name string, dsn string, readonly bool) {

	if pool.main == nil {
		pool.main = mysql.Open(dsn)
		return
	}

	if readonly {
		pool.replicas = append(pool.replicas, conn{dialector: mysql.Open(dsn), name: name})
	} else {
		pool.sources = append(pool.sources, conn{dialector: mysql.Open(dsn), name: name})
	}
}

// Connect connect db
func (pool *Pool) Connect() {
	db, err := gorm.Open(pool.main, &gorm.Config{})
	if err != nil {
		exception.Err(err, 500).
			Ctx(json.M{"pool": pool.main}).
			Throw()
	}
	pool.db = db
	var resolver *dbresolver.DBResolver
	if len(pool.sources) > 0 {
		sources := []gorm.Dialector{}
		names := []interface{}{}
		arr.Pluck(pool.sources, func(c conn) gorm.Dialector { return c.dialector }, &sources)
		arr.Pluck(pool.sources, func(c conn) interface{} { return c.name }, &names)
		resolver = dbresolver.Register(dbresolver.Config{Sources: sources}, names...)
	}
	if len(pool.replicas) > 0 {
		replicas := []gorm.Dialector{}
		names := []interface{}{}
		arr.Pluck(pool.replicas, func(c conn) gorm.Dialector { return c.dialector }, &replicas)
		arr.Pluck(pool.replicas, func(c conn) interface{} { return c.name }, &names)
		if resolver != nil {
			resolver.Register(dbresolver.Config{Replicas: replicas}, names...)
		} else {
			resolver = dbresolver.Register(dbresolver.Config{Replicas: replicas}, names...)
		}
	}
	if resolver != nil {
		pool.db.Use(resolver)
	}
}

// DB Select some connection
func (pool *Pool) DB() *gorm.DB {
	return pool.db
}

// Select Select the connection with given name
func (pool *Pool) Select(name string) *gorm.DB {
	return pool.db.Clauses(dbresolver.Use(name))
}

// Read Select the read connection
func (pool *Pool) Read() *gorm.DB {
	return pool.db.Clauses(dbresolver.Read)
}

// Write Select the write connection
func (pool *Pool) Write() *gorm.DB {
	return pool.db.Clauses(dbresolver.Write)
}
