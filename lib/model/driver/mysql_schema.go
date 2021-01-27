package driver

// Columns Get the Columns
func (m *MySQL) Columns() {
	println(m.DSN + " Columns")
}

// CreateColumn Create a column
func (m *MySQL) CreateColumn() {}

// UpsertColumn Update a column if not exist create it
func (m *MySQL) UpsertColumn() {}

// DeleteColumn Delete a column
func (m *MySQL) DeleteColumn() {}

// Relations Get the relations of the model
func (m *MySQL) Relations() {}

// CreateRelation Create a relation
func (m *MySQL) CreateRelation() {}

// UpsertRelation  Update a relation if not exist create it
func (m *MySQL) UpsertRelation() {}

// DeleteRelation Delete a relation
func (m *MySQL) DeleteRelation() {}

// Upgrade Upgrade the model struct
func (m *MySQL) Upgrade() {}
