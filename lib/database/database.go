package database

import (
	"fmt"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/json"
)

// UseDefault Get or create the default connections
func UseDefault() *Pool {
	name, has := config.Setting.Default["database"]
	if !has {
		exception.New("the default database config does not set!", 404).
			Ctx(json.M{"Default": config.Setting.Default}).
			Throw()
	}
	return Use(name)
}

// Use Get or create connections from the config
func Use(name string) *Pool {
	settings, has := config.Setting.Database[name]
	if !has {
		exception.New("Database config not found!", 404).
			Ctx(json.M{"name": name, "Database": config.Setting.Database}).
			Throw()
	}
	return UseSetting(settings...)
}

// UseSetting Get or create connections from given settings
func UseSetting(settings ...config.Database) *Pool {
	fmt.Printf("settings: %#v", settings)
	return &Pool{}
}

// Select Select one connection
func (conns *Pool) Select(readonly ...bool) {

}
