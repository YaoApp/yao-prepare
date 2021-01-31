package storage

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/database"
)

func init() {
	config.Setting = config.Load(".yset", path.Join(config.PWD(), "/../.."))
	database.DB = database.UseDefault().DB()
}

func TestUseDefault(t *testing.T) {
	fs := UseDefault()
	err := afero.WriteFile(fs, "/test.go", []byte("hello world"), os.ModePerm)
	fmt.Printf("err: %#v", err)
}
