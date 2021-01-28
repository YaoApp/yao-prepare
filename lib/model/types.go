package model

import "github.com/yaoapp/yao/lib/model/driver"

// Model the model query methods
type Model interface {
	GetOption() *driver.Option
	Query()
	Create()
	Upsert()
	Update()
	Delete()
	Exists()
	Import()
	Export()
	ORM()
}

// Schema the model schema methods
type Schema interface {
	Columns()
	CreateColumn()
	UpsertColumn()
	DeleteColumn()
	Relations()
	CreateRelation()
	UpsertRelation()
	DeleteRelation()
	Upgrade()
}
