package database

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
)

func TestUse(t *testing.T) {
	config.Setting = config.Load("config.json", path.Join(config.PWD(), "/../../config"))
	name := config.Setting.Default["database"]
	conns := Use(name)
	conns.Select()
}

func TestUseSetting(t *testing.T) {
	primary := config.Database{
		Driver:   "Mysql",
		DSN:      "root:123456@tcp(192.168.31.119:3306)/xiang?charset=utf8mb4&parseTime=True&loc=Local",
		Readonly: false,
	}
	Secondary := config.Database{
		Driver:   "Mysql",
		DSN:      "xiang:123456@tcp(192.168.31.119:3306)/xiang?charset=utf8mb4&parseTime=True&loc=Local",
		Readonly: true,
	}
	conns := UseSetting(primary, Secondary)
	conns.Select()
}
