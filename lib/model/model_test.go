package model

import (
	"testing"
)

func TestNew(t *testing.T) {
	user := New("User")
	user.Query()
}

func TestNewSchema(t *testing.T) {
	sch := NewSchema("User")
	sch.Columns()
}

// func TestLoadOption(t *testing.T) {
// 	option := LoadOption("User", []string{"cache", "file"})
// 	fmt.Printf("option: %#v\n", option)
// }
