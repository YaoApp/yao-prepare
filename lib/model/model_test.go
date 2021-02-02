package model

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/database"
	"github.com/yaoapp/yao/lib/yms"
)

func init() {
	config.Setting = config.Load(".yset", path.Join(config.PWD(), "/../.."))
	yms.Load(path.Join(config.PWD(), "/../../lib/yms/assets"), "system")
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
