package yms

import (
	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/t"
)

// EachField walk each EachField
func (file *File) EachField(cb func(field *Field)) {
	for i := range file.Fields {
		cb(&file.Fields[i])
	}
}

// parse file
func (file *File) parse() {
	if file.parsed {
		return
	}
	// checking the format and set field mapping
	file.FieldsMap = map[string]*Field{}
	for idx, field := range file.Fields {
		if field.Name == "" && field.Field == "" {
			exception.New("YMS file format error. (the field name is required)", 500).
				Ctx(t.M{"index": idx, "field": field, "file": file.path, "namespace": file.namespace}).
				Throw()
		}
		if field.Field == "" {
			file.Fields[idx].parsed = true
		}

		// parse the import file first
		if field.From != "" {
			from := Get(file.namespace, field.From)
			if from != file {
				from.parse()
			}
		}

		if field.Name != "" {
			file.FieldsMap[field.Name] = &file.Fields[idx]
		} else {
			file.FieldsMap[field.Field] = &file.Fields[idx]
		}
	}

	// parse each field
	for _, field := range file.FieldsMap {
		field.parse(file)
	}
	file.parsed = true
}
