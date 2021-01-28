package model

import (
	"strings"

	"github.com/yaoapp/yao/lib/exception"

	"github.com/gobuffalo/packr"
	"github.com/yaoapp/yao/lib/json"
	"github.com/yaoapp/yao/lib/model/driver"
)

var loadOptionMethods = map[string]func(string) *driver.Option{
	"code":  loadOptionFromCode,
	"cache": loadOptionFromCache,
	"file":  loadOptionFromFile,
	"db":    loadOptionFromDB,
}

//New Create a model instance.
//name: the model name.
//loadfrom: code/cache/file/db, default is code.
//
//examples:
//  user  := model.New("user") // Load from the models folder
//  order := model.New("/some/dir/order", "file", "db")  // Load from a given path, if not exists, load from DB.
func New(name string, loadfrom ...string) Model {
	if len(loadfrom) == 0 {
		loadfrom = []string{"code"}
	}
	option := LoadOption(name, loadfrom)
	switch option.Engine.Storage {
	case "MySQL":
		return &driver.MySQL{DSN: name, Option: option}
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
// methods: code/cache/file/db
//
//examples:
//   // load the "user" model option from the cache first. if not exist load from file.
//   option := model.LoadOption("user", []string{"cache", "file"});
func LoadOption(name string, methods []string) *driver.Option {
	var option *driver.Option
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
func loadOptionFromCode(name string) *driver.Option {
	option := &driver.Option{}
	filename := strings.ToLower(name) + ".json"
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
func loadOptionFromCache(name string) *driver.Option {
	return nil
}

// loadOptionFromFile load a model option from the file
func loadOptionFromFile(filename string) *driver.Option {
	option := &driver.Option{}
	json.DecodeFile(filename, option)
	return option
}

// loadOptionFromDB load a model option from the database
func loadOptionFromDB(name string) *driver.Option {
	return nil
}
