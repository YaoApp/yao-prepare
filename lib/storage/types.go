package storage

import "github.com/spf13/afero"

// Fs is the filesystem interface.
type Fs struct {
	root string // The root path
	afero.Afero
}
