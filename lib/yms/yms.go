// Package yms The (Y)ao (M)odel (S)pecification
package yms

import (
	"os"
	"path/filepath"

	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/storage"
	"github.com/yaoapp/yao/lib/t"
)

// space The memory space for saving parse YMS file.
var space map[string]storage.Fs

func init() {
	space = map[string]storage.Fs{}
}

// Load load the YMS files into the memory space.
func Load(path string, namespace string) {

	if _, has := space[namespace]; has {
		exception.New("namespace has been used!", 500).
			Ctx(t.M{"path": path, namespace: namespace}).
			Throw()
	}

	space[namespace] = storage.MemFs(namespace)
	fs := storage.OsFs(path)
	fs.Walk("/", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".yms" {
			return nil
		}
		contents, err := fs.ReadFile(path)
		if err != nil {
			exception.Err(err, 500).
				Ctx(t.M{"path": path, namespace: namespace}).
				Throw()
		}
		space[namespace].WriteFile(path, contents, 0644)
		return nil
	})
}

// Open open the YMS file.
func Open(name string, filename string, namespace string) *File {
	return &File{}
}
