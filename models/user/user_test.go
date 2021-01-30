package user

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/database"
)

func init() {
	config.Setting = config.Load(".yao.env", path.Join(config.PWD(), "/../.."))
	database.DB = database.UseDefault().DB()
}

func TestNew(t *testing.T) {
	user := New()
	user.Query()
	user.Login()
}
