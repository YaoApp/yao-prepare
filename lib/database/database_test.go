package database

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/json"
)

func init() {
	config.Setting = config.Load(".yset", path.Join(config.PWD(), "/../.."))
	DB = UseDefault().DB()
}

func TestUseDefault(t *testing.T) {
	apps := []map[string]interface{}{}
	DB.Debug().Table("app").Limit(2).Find(&apps)
	json.PrettyPrint(apps)
}

func TestUse(t *testing.T) {
	pool := Use("main")
	db := pool.Read()
	rows := []map[string]interface{}{}
	db.Debug().Raw("SELECT * FROM information_schema.processlist WHERE COMMAND='Query'").Find(&rows)
	json.PrettyPrint(rows)
}

func TestUseSetting(t *testing.T) {
	primary := config.Database{
		Name:     "primary",
		Driver:   "MySQL",
		DSN:      "root:123456@tcp(192.168.31.119:3306)/xiang?charset=utf8mb4&parseTime=True&loc=Local",
		Readonly: false,
	}
	Secondary := config.Database{
		Name:     "secondary",
		Driver:   "MySQL",
		DSN:      "xiang:123456@tcp(192.168.31.119:3306)/xiang?charset=utf8mb4&parseTime=True&loc=Local",
		Readonly: true,
	}
	pool := UseSetting(primary, Secondary)
	db := pool.Select("secondary")
	rows := []map[string]interface{}{}
	db.Debug().Raw("SELECT * FROM information_schema.processlist WHERE COMMAND='Query'").Find(&rows)
	json.PrettyPrint(rows)
}
