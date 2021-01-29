package database

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
)

func init() {
	config.Setting = config.Load("config.json", path.Join(config.PWD(), "/../../config"))
}

func TestUseDefault(t *testing.T) {
	conns := UseDefault()
	conns.Select()
}

func TestUse(t *testing.T) {
	conns := Use("main")
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
