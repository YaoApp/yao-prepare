package model

import (
	"path"
	"strings"

	"github.com/yaoapp/yao/lib/json"
	"github.com/yaoapp/yao/lib/model/driver"
)

var loadOptionMethods = map[string]func(string) *Option{
	"global": loadOptionFromGlobal,
	"cache":  loadOptionFromCache,
	"file":   loadOptionFromFile,
	"db":     loadOptionFromDB,
}

//New Create a Model instance
func New(name string, engine string) Model {
	switch engine {
	case "MySQL":
		var m Model = &driver.MySQL{DSN: name}
		return m
	}
	return nil
}

// NewSchema Create a model schema instance
func NewSchema(name string, engine string) Schema {
	switch engine {
	case "MySQL":
		var m Schema = &driver.MySQL{DSN: name}
		return m
	}
	return nil
}

// LoadOption load a model option.
// name: the name of model.
// methods: global/cache/file/db
//
// LoadOption("user", []string{"cache", "file"});
// load the "user" model option from the cache first. if not exist load from file.
func LoadOption(name string, methods []string) *Option {
	var option *Option
	for _, method := range methods {
		load, has := loadOptionMethods[method]
		if has {
			option = load(name)
		}
		if option != nil {
			return option
		}
	}
	return nil
}

// loadOptionFromGlobal load a model option from the global variable
func loadOptionFromGlobal(name string) *Option {
	return nil
}

// loadOptionFromCache load a model option from the cache
func loadOptionFromCache(name string) *Option {
	return nil
}

// loadOptionFromFile load a model option from the file
func loadOptionFromFile(name string) *Option {
	option := &Option{}
	filename := path.Join("/Users/max/Code/yao/yao/models", strings.ToLower(name)+".json")
	json.DecodeFile(filename, option)
	return option
}

// loadOptionFromDB load a model option from the database
func loadOptionFromDB(name string) *Option {
	return nil
}
