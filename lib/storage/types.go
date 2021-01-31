package storage

import (
	"net/http"

	"github.com/spf13/afero"
)

// Fs is the filesystem interface.
type Fs struct {
	root string // The root path
	afero.Afero
}

// BinFs the http file system interface
type BinFs struct {
	http.FileSystem
}
