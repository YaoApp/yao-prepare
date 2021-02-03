package model

import (
	"path/filepath"
	"strings"

	"github.com/yaoapp/yao/lib/database"
	"github.com/yaoapp/yao/lib/model/engine/db"
	"github.com/yaoapp/yao/lib/yms"
	"gorm.io/gorm"
)

//New Create a model instance.
//
//examples:
//  user  := model.New("user")
func New(name string) Model {
	return NewWithDB(database.DB, name)
}

//NewSchema Create a model schema instance.
//
//examples:
//  userSchema  := model.NewSchema("user")
func NewSchema(name string) Schema {
	return NewSchemaWithDB(database.DB, name)
}

// NewWithDB Create a model instance and bind connection .
func NewWithDB(gormDB *gorm.DB, name string) Model {
	filename := YMSFilePath(name)
	file := yms.Get("system", filename)
	switch file.Engine {
	case "db", "orm", "database", "gorm":
		var m Model = &db.Engine{DB: gormDB, File: file}
		return m
	}
	return nil
}

// NewSchemaWithDB Create a model instance and bind connection .
func NewSchemaWithDB(gormDB *gorm.DB, name string) Schema {
	filename := YMSFilePath(name)
	file := yms.Get("system", filename)
	switch file.Engine {
	case "db", "orm", "database", "gorm":
		var schema Schema = &db.Engine{DB: gormDB, File: file}
		return schema
	}
	return nil
}

// YMSFilePath Get the YMS file path
func YMSFilePath(name string) string {
	return strings.ToLower(filepath.Join("/", name, "model.yms"))
}
