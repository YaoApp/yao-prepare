package database

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/json"
)

func init() {
	config.Setting = config.Load(".yao.env", path.Join(config.PWD(), "/../.."))
}

func TestUseDefault(t *testing.T) {
	pool := UseDefault()
	db := pool.DB()
	apps := []map[string]interface{}{}
	db.Debug().Table("app").Limit(2).Find(&apps)
	json.PrettyPrint(apps)
}

func TestUse(t *testing.T) {
	conns := Use("main")
	conns.DB()
}

func TestUseSetting(t *testing.T) {
	primary := config.Database{
		Name:     "primary",
		Driver:   "Mysql",
		DSN:      "root:123456@tcp(192.168.31.119:3306)/xiang?charset=utf8mb4&parseTime=True&loc=Local",
		Readonly: false,
	}
	Secondary := config.Database{
		Name:     "secondary",
		Driver:   "Mysql",
		DSN:      "xiang:123456@tcp(192.168.31.119:3306)/xiang?charset=utf8mb4&parseTime=True&loc=Local",
		Readonly: true,
	}
	conns := UseSetting(primary, Secondary)
	conns.DB()
}
