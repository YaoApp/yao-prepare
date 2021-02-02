package engine

import (
	"fmt"

	"github.com/yaoapp/yao/lib/yms"
	"gorm.io/gorm"
)

// ORM ORM driver
type ORM struct {
	DB   *gorm.DB
	File *yms.File
}

// ORM Return the ORM DB instance
func (m *ORM) ORM() *gorm.DB {
	return m.DB
}

// Query query by params
func (m *ORM) Query() {
	fmt.Printf("Query table: %s\n", m.File.Table.Name)
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
