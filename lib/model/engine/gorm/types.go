package gorm

import "github.com/yaoapp/yao/lib/yms"

// Table the database table
type Table struct {
	yms.Table
	Fields  []yms.Field
	Columns []*Column
}

// Column the database table column
type Column struct {
	Comment  string
	Name     string
	Type     string
	Length   *int
	Args     interface{}
	Default  interface{}
	Nullable *bool
	Unsigned *bool
}
