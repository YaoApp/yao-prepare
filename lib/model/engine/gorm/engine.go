package gorm

import (
	"fmt"

	"github.com/yaoapp/yao/lib/yms"
	"gorm.io/gorm"
)

// Engine Engine driver
type Engine struct {
	DB   *gorm.DB
	File *yms.File
}

// GORM Return the GORM DB instance
func (m *Engine) GORM() *gorm.DB {
	return m.DB
}

// Query query by params
func (m *Engine) Query() {
	fmt.Printf("Query table: %s\n", m.File.Table.Name)
}

// Create Create
func (m *Engine) Create() {}

// Upsert Update if not exist create it
func (m *Engine) Upsert() {}

// Update Update the existing record
func (m *Engine) Update() {}

// Delete Delete the existing record
func (m *Engine) Delete() {}

// Exists Check the record is exists.
func (m *Engine) Exists() {}

// Import Import by given conditions
func (m *Engine) Import() {}

// Export Export by given conditions
func (m *Engine) Export() {}
