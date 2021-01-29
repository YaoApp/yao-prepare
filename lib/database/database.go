package database

import (
	"fmt"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/json"
)

// Use Get or create connections from config
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
