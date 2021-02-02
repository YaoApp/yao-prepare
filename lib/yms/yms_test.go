package yms

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/database"
	"github.com/yaoapp/yao/lib/json"
)

func init() {
	config.Setting = config.Load(".yset", path.Join(config.PWD(), "/../.."))
	database.DB = database.UseDefault().DB()
}

func TestLoad(t *testing.T) {
	Load(path.Join(config.PWD(), "assets"), "test")
	MustHave("test")
	MustHaveFile("test",
		"/guider/model.yms",
		"/user/model.yms",
		"/ec/error.yms",
		"/ec/user.yms",
	)
}

func TestNamespaces(t *testing.T) {
	namespaces := Namespaces()
	json.PrettyPrint(namespaces)
}

func TestFiles(t *testing.T) {
	files := Files("test")
	json.PrettyPrint(files)
}
