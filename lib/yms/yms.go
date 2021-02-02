// Package yms The (Y)ao (M)odel (S)pecification
package yms

import (
	"os"
	"path/filepath"

	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/json"
	"github.com/yaoapp/yao/lib/maps"
	"github.com/yaoapp/yao/lib/storage"
	"github.com/yaoapp/yao/lib/t"
)

// files The memory space for saving YMS file object.
var files map[string]map[string]*File

func init() {
	files = map[string]map[string]*File{}
}

// Load load the YMS files into the memory space.
func Load(path string, namespace string) {

	if _, has := files[namespace]; has {
		exception.New("namespace has been used!", 500).
			Ctx(t.M{"path": path, namespace: namespace}).
			Throw()
	}
	files[namespace] = map[string]*File{}

	fs := storage.OsFs(path)
	if exists, err := fs.Exists("/"); err != nil || !exists {
		exception.New(path+" does not exists", 500).
			Ctx(t.M{"path": path, namespace: namespace}).
			Throw()
	}

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

		file := &File{}
		json.DecodeBytes(contents, file)
		file.namespace = namespace
		file.parsed = false
		file.path = path
		files[namespace][path] = file
		return nil
	})

	// parse the files
	for _, file := range files[namespace] {
		file.parse()
	}
}

// Get get the YMS File
func Get(namespace string, filename string) *File {
	MustHaveFile(namespace, filename)
	return files[namespace][filename]
}

// Files get all files name of the namespace
func Files(namespace string) []string {
	MustHave(namespace)
	return maps.Keys(files[namespace])
}

// Namespaces get all loaded namespaces
func Namespaces() []string {
	return maps.Keys(files)
}

// MustHave the given namespace must be loaded
func MustHave(namespace string) {
	if _, has := files[namespace]; !has {
		exception.New("namespace does not exists", 500).
			Ctx(t.M{"namespace": namespace}).
			Throw()
	}
}

// MustHaveFile the given file must be loaded
func MustHaveFile(namespace string, filenames ...string) {
	MustHave(namespace)
	for _, filename := range filenames {
		_, has := files[namespace][filename]
		if !has {
			exception.New(filename+" does not exists", 500).
				Ctx(t.M{"namespace": namespace, "filename": filename}).
				Throw()
		}
	}
}
