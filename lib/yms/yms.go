// Package yms The (Y)ao (M)odel (S)pecification
package yms

import (
	"os"
	"path/filepath"

	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/json"
	"github.com/yaoapp/yao/lib/storage"
	"github.com/yaoapp/yao/lib/t"
)

// files The memory space for saving YMS file object.
var files map[string]map[string]*File

func init() {
	files = map[string]map[string]*File{}
}

// Load load the YMS files into the memory space.
func Load(path string, namespace string) {

	if _, has := files[namespace]; has {
		exception.New("namespace has been used!", 500).
			Ctx(t.M{"path": path, namespace: namespace}).
			Throw()
	}
	files[namespace] = map[string]*File{}
	fs := storage.OsFs(path)
	fs.Walk("/", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".yms" {
			return nil
		}
		contents, err := fs.ReadFile(path)
		if err != nil {
			exception.Err(err, 500).
				Ctx(t.M{"path": path, namespace: namespace}).
				Throw()
		}

		file := &File{}
		json.DecodeBytes(contents, file)
		file.namespace = namespace
		file.parsed = false
		file.path = path
		files[namespace][path] = file
		return nil
	})

	// parse the files
	for _, file := range files[namespace] {
		file.parse()
	}
}

// Get get the YMS File
func Get(namespace string, filename string) *File {
	if _, has := files[namespace]; !has {
		exception.New("namespace does not exists", 500).
			Ctx(t.M{"namespace": namespace, "filename": filename}).
			Throw()
	}

	file, has := files[namespace][filename]
	if !has {
		exception.New(filename+" does not exists", 500).
			Ctx(t.M{"namespace": namespace, "filename": filename}).
			Throw()
	}
	return file
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
