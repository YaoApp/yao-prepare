package yms

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/yao/lib/arr"
	"github.com/yaoapp/yao/test"
)

func init() {
	test.Use("unit")
}

func TestLoad(t *testing.T) {
	Load(path.Join(test.Root(), "/lib/yms/assets"), "test")
	MustHave("test")
	MustHaveFile("test",
		"/guider/model.yms",
		"/user/model.yms",
		"/ec/error.yms",
		"/ec/user.yms",
	)
	assert.True(t, true)
}

func TestNamespaces(t *testing.T) {
	namespaces := Namespaces()
	shouldbe := "test"
	assert.True(t, arr.Have(namespaces, "test"), `the loaded namespaces should be %#v`, shouldbe)
}

func TestFiles(t *testing.T) {
	files := Files("test")
	shouldbe := []string{
		"/ec/error.yms",
		"/ec/user.yms",
		"/guider/model.yms",
		"/user/model.yms",
	}
	assert.True(t, arr.HaveMany(files, shouldbe), `the loaded files should be %#v`, shouldbe)
}
