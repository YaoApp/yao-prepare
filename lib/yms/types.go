package yms

// Field the field description struct
type Field struct {
	Comment  string      `json:"comment,omitempty"`
	Name     string      `json:"name"`
	Type     string      `json:"string"`
	Args     interface{} `json:"args,omitempty"`
	Example  interface{} `json:"example,omitempty"`
	Nullable bool        `json:"nullable,omitempty"`
	Generate string      `json:"generate,omitempty"` // Increment, UUID,...
	Encoder  string      `json:"encoder,omitempty"`  // AES-256, AES-128, PASSWORD-HASH, ...
	Decoder  string      `json:"decoder,omitempty"`  // AES-256, AES-128, ...
	Extra    struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Type        string `json:"type,omitempty"`
		Pattern     string `json:"pattern,omitempty"`
	} `json:"extra,omitempty"`
	Field string `json:"field,omitempty"` // Extend from the parent field
	From  string `json:"from,omitempty"`  // Extend from which structure
}

// Search the search option struct
type Search struct {
	Field  string   `json:"field,omitempty"`
	Fields []string `json:"fields,omitempty"`
	Type   string   `json:"string"` // primary,unique,index,match
}

// Table the model mapping table in DB
type Table struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name"`
	Engine  string `json:"engine,omitempty"` // InnoDB,MyISAM ( MySQL Only )
}

// File the YMS file.
type File struct {
	Name   string   `json:"name"`
	Table  Table    `json:"table,omitempty"`
	Engine string   `json:"engine,omitempty"`
	Fields []Field  `json:"fields"`
	Search []Search `json:"search"`
}
