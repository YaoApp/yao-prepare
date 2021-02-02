package test

import (
	"os"
	"path/filepath"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/database"
	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/storage"
	"github.com/yaoapp/yao/lib/sys"
	"github.com/yaoapp/yao/lib/t"
)

var root string = ""

// Use set the testing environment
func Use(name string) {
	env := os.Getenv("YAO_TEST")
	if env != "" {
		name = env
	}
	configfile := "." + name + ".test.yset"
	if root == "" {
		root = findPath(configfile, sys.PWD())
	}

	// init configture
	config.Setting = config.Load(configfile, root)
	database.DB = database.UseDefault().DB()
}

// Root Get the root path
func Root() string {
	return root
}

func findPath(name string, path string) string {
	filename := filepath.Join(path, name)
	fs := storage.OsFs("/")
	if exists, _ := fs.Exists(filename); exists {
		return path
	} else if path != "/" {
		path = filepath.Dir(path)
		return findPath(name, path)
	}
	exception.New("the testing config file not found", 404).
		Ctx(t.M{"name": name}).
		Throw()
	return ""
}
