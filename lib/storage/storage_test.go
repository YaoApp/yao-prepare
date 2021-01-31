package storage

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/database"
)

func init() {
	config.Setting = config.Load(".yset", path.Join(config.PWD(), "/../.."))
	database.DB = database.UseDefault().DB()
}

func TestUseDefault(t *testing.T) {
	fs := UseDefault()
	err := fs.WriteFile("/test.md", []byte("hello world"), os.ModePerm)
	fmt.Printf("err: %#v", err)
}

func TestMemFs(t *testing.T) {
	fs := MemFs("testing")
	err := fs.WriteFile("/test.md", []byte("hello world"), os.ModePerm)
	fmt.Printf("err: %#v\n\n", err)

	if has, err := fs.Exists("/test.md"); err == nil && has {
		fmt.Printf("/test.md has found\n\n")
	}

	content, err := fs.ReadFile("/test.md")
	fmt.Printf("%s ERR:%#v\n", content, err)
}
