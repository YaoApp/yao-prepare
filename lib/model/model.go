package model

import (
	"strings"

	"github.com/yaoapp/yao/lib/database"
	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/model/engine"
	"gorm.io/gorm"

	"github.com/gobuffalo/packr"
	"github.com/yaoapp/yao/lib/json"
)

var loadOptionMethods = map[string]func(string) *engine.Option{
	"code":  loadOptionFromCode,
	"cache": loadOptionFromCache,
	"file":  loadOptionFromFile,
	"db":    loadOptionFromDB,
}

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
	option := LoadOption(name, []string{"code"})
	switch option.Engine {
	case "orm":
		var m Model = &engine.ORM{DB: db, Option: option}
		return m
	}
	return nil
}

// NewSchemaWithDB Create a model instance and bind connection .
func NewSchemaWithDB(db *gorm.DB, name string) Schema {
	option := LoadOption(name, []string{"code"})
	switch option.Engine {
	case "orm":
		var schema Schema = &engine.ORM{DB: db, Option: option}
		return schema
	}
	return nil
}

// LoadOption load a model option.
// name: the name of model.
// methods: code/cache/file/db
//
//examples:
//   // load the "user" model option from the cache first. if not exist load from file.
//   option := model.LoadOption("user", []string{"cache", "file"});
func LoadOption(name string, methods []string) *engine.Option {
	var option *engine.Option
	for _, method := range methods {
		load, has := loadOptionMethods[method]
		if has {
			option = load(name)
		}
		if option != nil {
			return option
		}
	}
	exception.New(name+" not found!", 404).
		Ctx(json.M{"name": name, "methods": methods}).
		Throw()
	return nil
}

// loadOptionFromCode load a model option from /models folder .
func loadOptionFromCode(name string) *engine.Option {
	option := &engine.Option{}
	filename := strings.ToLower(name) + "/model.json"
	box := packr.NewBox("../../models")
	content, err := box.FindString(filename)
	if err != nil {
		exception.Err(err, 500).
			Ctx(json.M{"filename": filename}).
			Throw()
	}
	json.Decode(content, option)
	return option
}

// loadOptionFromCache load a model option from the cache
func loadOptionFromCache(name string) *engine.Option {
	return nil
}

// loadOptionFromFile load a model option from the file
func loadOptionFromFile(filename string) *engine.Option {
	option := &engine.Option{}
	json.DecodeFile(filename, option)
	return option
}

// loadOptionFromDB load a model option from the database
func loadOptionFromDB(name string) *engine.Option {
	return nil
}
