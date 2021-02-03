package gorm

import (
	"github.com/yaoapp/yao/lib/json"
	"github.com/yaoapp/yao/lib/t"
)

// Schema get the databae table schema
func (table *Table) Schema() {
	json.PrettyPrint(table.Columns)
}

// parseColumns parse fields to columns
func (table *Table) parseColumns() {
	columns := []*Column{}
	for _, field := range table.Fields {
		typ := field.Type
		if typ == "" {
			typ = "string"
		}
		nullable := field.Nullable
		if nullable == nil {
			nullable = t.Bool(false)
		}

		length := columnLength(typ, field.Args)
		if length == nil {
			length = columnLengthDefault(typ)
		}

		unsigned := field.Unsigned
		if unsigned == nil {
			unsigned = t.Bool(false)
		}
		column := &Column{
			Name:     field.Name,
			Type:     typ,
			Default:  field.Default,
			Comment:  field.Comment,
			Length:   length,
			Args:     field.Args,
			Nullable: nullable,
			Unsigned: unsigned,
		}
		columns = append(columns, column)
	}
	table.Columns = columns
}
