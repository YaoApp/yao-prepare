package model

import (
	"github.com/yaoapp/yao/lib/model/engine"
	"gorm.io/gorm"
)

// Model the model query methods
type Model interface {
	GetOption() *engine.Option
	GetDB() *gorm.DB
	Query()
	Create()
	Upsert()
	Update()
	Delete()
	Exists()
	Import()
	Export()
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
