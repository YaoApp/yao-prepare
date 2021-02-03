package db

import (
	"fmt"
)

// Columns Get the Columns
func (m *Engine) Columns() {
	fmt.Printf("Columns table: %s\n", m.File.Table.Name)
}

// CreateColumn Create a column
func (m *Engine) CreateColumn() {}

// UpsertColumn Update a column if not exist create it
func (m *Engine) UpsertColumn() {}

// DeleteColumn Delete a column
func (m *Engine) DeleteColumn() {}

// Relations Get the relations of the model
func (m *Engine) Relations() {}

// CreateRelation Create a relation
func (m *Engine) CreateRelation() {}

// UpsertRelation  Update a relation if not exist create it
func (m *Engine) UpsertRelation() {}

// DeleteRelation Delete a relation
func (m *Engine) DeleteRelation() {}

// Upgrade Upgrade the model struct
func (m *Engine) Upgrade() {
	m.Table().Schema()
}

// Table return the databae table pointer
func (m *Engine) Table() *Table {
	table := &Table{
		Table:  m.File.Table,
		Fields: m.File.Fields,
	}
	table.parseColumns()
	return table
}
