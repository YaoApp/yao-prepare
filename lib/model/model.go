package model

import (
	"path/filepath"
	"strings"

	"github.com/yaoapp/yao/lib/database"
	"github.com/yaoapp/yao/lib/model/engine"
	"github.com/yaoapp/yao/lib/yms"
	"gorm.io/gorm"
)

//New Create a model instance.
//
//examples:
//  user  := model.New("user") // Load from the models folder
func New(name string) Model {
	return NewWithDB(database.DB, name)
}

//NewSchema Create a model schema instance.
//
//examples:
//  userSchema  := model.NewSchema("user") // Load from the models folder
func NewSchema(name string) Schema {
	return NewSchemaWithDB(database.DB, name)
}

// NewWithDB Create a model instance and bind connection .
func NewWithDB(db *gorm.DB, name string) Model {
	filename := strings.ToLower(filepath.Join("/", name, "model.yms"))
	file := yms.Get("system", filename)
	switch file.Engine {
	case "orm":
		var m Model = &engine.ORM{DB: db, File: file}
		return m
	}
	return nil
}

// NewSchemaWithDB Create a model instance and bind connection .
func NewSchemaWithDB(db *gorm.DB, name string) Schema {
	filename := strings.ToLower(filepath.Join("/", name, "model.yms"))
	file := yms.Get("system", filename)
	switch file.Engine {
	case "orm":
		var schema Schema = &engine.ORM{DB: db, File: file}
		return schema
	}
	return nil
}
