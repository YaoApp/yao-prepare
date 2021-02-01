package yms

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

func TestLoad(t *testing.T) {
	Load(path.Join(config.PWD(), "assets"), "test")
}
