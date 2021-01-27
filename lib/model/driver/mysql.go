package driver

// MySQL MySQL driver
type MySQL struct {
	DSN string
}

// Query query by params
func (m *MySQL) Query() {
	println(m.DSN + " Query")
}

// Create Create
func (m *MySQL) Create() {}

// Upsert Update if not exist create it
func (m *MySQL) Upsert() {}

// Update Update the existing record
func (m *MySQL) Update() {}

// Delete Delete the existing record
func (m *MySQL) Delete() {}

// Exists Check the record is exists.
func (m *MySQL) Exists() {}

// Import Import by given conditions
func (m *MySQL) Import() {}

// Export Export by given conditions
func (m *MySQL) Export() {}

// ORM Get the ORM instance
func (m *MySQL) ORM() {}
