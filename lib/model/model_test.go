package model

import (
	"path"
	"testing"

	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/yms"
	"github.com/yaoapp/yao/test"
)

func init() {
	test.Use("unit")
	yms.Load(path.Join(test.Root(), "/lib/yms/assets"), "system")
}

func TestUpgrade(t *testing.T) {
	defer exception.CatchDebug()
	schema := NewSchema("User")
	schema.Upgrade()
}

// func TestNew(t *testing.T) {
// 	user := New("User")
// 	user.Query()
// }
