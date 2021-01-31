package model

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/database"
)

func init() {
	config.Setting = config.Load(".yset", path.Join(config.PWD(), "/../.."))
	database.DB = database.UseDefault().DB()
}

func TestNew(t *testing.T) {
	user := New("User")
	user.Query()
}

func TestNewSchema(t *testing.T) {
	sch := NewSchema("User")
	sch.Columns()
}

// func TestLoadOption(t *testing.T) {
// 	option := LoadOption("User", []string{"cache", "file"})
// 	fmt.Printf("option: %#v\n", option)
// }
