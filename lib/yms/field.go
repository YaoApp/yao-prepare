package yms

import (
	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/t"
)

// parse field
func (field *Field) parse(file *File) {
	if field.parsed {
		return
	}

	name := field.Name
	if name == "" {
		name = field.Field
	}

	var parent *Field
	var has bool

	// get the import field
	if field.From == "" { // at the same YMS file.  {"comment": "真名","field": "nickname","name": "name"}
		parent, has = file.FieldsMap[field.Field]
		if !has {
			exception.New(field.Field+" does not exists", 500).
				Ctx(t.M{"field": field, "file": field.From, "namespace": file.namespace}).
				Throw()
		}
		if field == parent {
			exception.New(name+" can't import self field", 500).
				Ctx(t.M{"field": field, "file": field.From, "namespace": file.namespace}).
				Throw()
		}
		parent.parse(file)
	} else { // at the different YMS file. {"field": "mobile","from": "/ec/user.yms"},
		from := Get(file.namespace, field.From)
		parent, has = from.FieldsMap[field.Field]
		if !has {
			exception.New(field.From+":"+field.Field+" does not exists", 500).
				Ctx(t.M{"field": field, "file": field.From, "namespace": file.namespace}).
				Throw()
		}
		if field == parent {
			exception.New(name+" can't import self field", 500).
				Ctx(t.M{"field": field, "file": field.From, "namespace": file.namespace}).
				Throw()
		}
		parent.parse(file)
	}

	// merge the import field info to field
	if parent != nil {
		field.merge(parent)
	}
	field.parsed = true
	// fmt.Printf("\tName: %s \tComment:%s \tField: %s \tFrom:%s| \t parent: %s %s\n", field.Name, field.Comment, field.Field, field.From, parent.Comment, parent.Name)
}

// merge field  ( ⚠️ this method should optimization)
func (field *Field) merge(fields ...*Field) {
	for _, new := range fields {
		if field.Comment == "" && new.Comment != "" {
			field.Comment = new.Comment
		}
		if field.Name == "" && new.Name != "" {
			field.Name = new.Name
		}
		if field.Type == "" && new.Type != "" {
			field.Type = new.Type
		}
		if field.Args == nil && new.Args != nil {
			field.Args = new.Args
		}
		if field.Example == nil && new.Example != nil {
			field.Example = new.Example
		}
		if field.Nullable == nil && new.Nullable != nil {
			field.Nullable = new.Nullable
		}
		if field.Generate == "" && new.Generate != "" {
			field.Generate = new.Generate
		}
		if field.Encoder == "" && new.Encoder != "" {
			field.Encoder = new.Encoder
		}
		if field.Decoder == "" && new.Decoder != "" {
			field.Decoder = new.Decoder
		}
		if new.Extra != nil {
			if field.Extra == nil {
				field.Extra = new.Extra
				continue
			}
			if field.Extra.Title == "" && new.Extra.Title != "" {
				field.Extra.Title = new.Extra.Title
			}
			if field.Extra.Description == "" && new.Extra.Description != "" {
				field.Extra.Description = new.Extra.Description
			}
			if field.Extra.Type == "" && new.Extra.Type != "" {
				field.Extra.Type = new.Extra.Type
			}
			if field.Extra.Pattern == "" && new.Extra.Pattern != "" {
				field.Extra.Pattern = new.Extra.Pattern
			}
		}
	}
}
