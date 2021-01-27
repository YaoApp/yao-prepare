package model

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	m := New("Hello", "MySQL")
	m.Query()
}

func TestNewSchema(t *testing.T) {
	sch := NewSchema("Hello", "MySQL")
	sch.Columns()
}

func TestLoadOption(t *testing.T) {
	option := LoadOption("User", []string{"cache", "file"})
	fmt.Printf("option: %#v\n", option)
}
