package engine

import "fmt"

// Columns Get the Columns
func (m *ORM) Columns() {
	fmt.Printf("Columns Engine is: %s\n", m.Option.Engine)
}

// CreateColumn Create a column
func (m *ORM) CreateColumn() {}

// UpsertColumn Update a column if not exist create it
func (m *ORM) UpsertColumn() {}

// DeleteColumn Delete a column
func (m *ORM) DeleteColumn() {}

// Relations Get the relations of the model
func (m *ORM) Relations() {}

// CreateRelation Create a relation
func (m *ORM) CreateRelation() {}

// UpsertRelation  Update a relation if not exist create it
func (m *ORM) UpsertRelation() {}

// DeleteRelation Delete a relation
func (m *ORM) DeleteRelation() {}

// Upgrade Upgrade the model struct
func (m *ORM) Upgrade() {}
