package driver

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
}

// SearchOption the search option struct
type SearchOption struct {
	Field  string   `json:"field,omitempty"`
	Fields []string `json:"fields,omitempty"`
	Type   string   `json:"string"` // primary,unique,index,match
}

// Option the model option
type Option struct {
	Name   string `json:"name"`
	DSN    string `json:"dsn,omitempty"`
	Engine struct {
		Storage string `json:"storage,omitempty"`
		Query   string `json:"query,omitempty"`
	}
	Fields []Field        `json:"fields"`
	Search []SearchOption `json:"search"`
}
