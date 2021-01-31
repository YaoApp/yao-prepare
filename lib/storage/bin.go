package storage

import (
	"io/ioutil"
	"net/http"
	"os"

	statikfs "github.com/rakyll/statik/fs"
	"github.com/yaoapp/yao/lib/exception"
	_ "github.com/yaoapp/yao/statik" // the static data
)

// Bin Return the Statik fs
func Bin() BinFs {
	fs, err := statikfs.New()
	if err != nil {
		exception.Err(err, 500).
			Throw()
	}
	return BinFs{
		FileSystem: fs,
	}
}

// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
func (fs BinFs) ReadFile(filename string) ([]byte, error) {
	r, err := fs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

// Readdir reads the directory named by dirname and returns
// a list directory entries.
func (fs BinFs) Readdir(dirname string) ([]os.FileInfo, error) {
	f, err := fs.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	return list, nil
}

// ListenAndServe start an http server
func (fs BinFs) ListenAndServe(addr string, path string, route string) {
	http.Handle(path, http.StripPrefix(route, http.FileServer(fs)))
	http.ListenAndServe(addr, nil)
}
