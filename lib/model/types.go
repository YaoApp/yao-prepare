package model

import (
	"gorm.io/gorm"
)

// Model the model query methods
type Model interface {
	Query()
	Create()
	Upsert()
	Update()
	Delete()
	Exists()
	Import()
	Export()
	GORM() *gorm.DB
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
	GORM() *gorm.DB
}
