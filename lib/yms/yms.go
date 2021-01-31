// Package yms The (Y)ao (M)odel (S)pecification
package yms

// LoadSYS load the system YMS files into the memory space.
func LoadSYS() {}

// LoadAPP load the application YMS files into the memory space.
func LoadAPP(name string) {}

// Load load the YMS files into the memory space.
func Load(path string, namespace string) {}

// OpenSYS open the system YMS file.
func OpenSYS(file string) *File {
	return Open("system", file, "system")
}

// OpenAPP open the application YMS file.
func OpenAPP(name string, file string) *File {
	return Open(name, file, "app")
}

// Open open the YMS file.
func Open(name string, file string, namespace string) *File {
	return &File{}
}

func parse() {}
