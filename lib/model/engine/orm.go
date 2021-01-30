package engine

import (
	"github.com/yaoapp/yao/lib/json"
	"gorm.io/gorm"
)

// ORM ORM driver
type ORM struct {
	DB     *gorm.DB
	Option *Option
}

// GetOption Get the model option
func (m *ORM) GetOption() *Option {
	return m.Option
}

// GetDB Get the DB instance
func (m *ORM) GetDB() *gorm.DB {
	return m.DB
}

// Query query by params
func (m *ORM) Query() {
	rows := []map[string]interface{}{}
	m.DB.Table(m.Option.Table.Name).Limit(20).Find(&rows)
	json.PrettyPrint(rows)
}

// Create Create
func (m *ORM) Create() {}

// Upsert Update if not exist create it
func (m *ORM) Upsert() {}

// Update Update the existing record
func (m *ORM) Update() {}

// Delete Delete the existing record
func (m *ORM) Delete() {}

// Exists Check the record is exists.
func (m *ORM) Exists() {}

// Import Import by given conditions
func (m *ORM) Import() {}

// Export Export by given conditions
func (m *ORM) Export() {}
