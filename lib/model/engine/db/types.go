package db

import (
	"github.com/yaoapp/yao/lib/yms"
	"gorm.io/gorm"
)

// Table the database table
type Table struct {
	yms.Table
	Fields  []yms.Field
	Columns []*Column
	DB      *gorm.DB
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
