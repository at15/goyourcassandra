package types

// Bookmark contains host, credential and template for common query
// TODO: should be able to read from .cassandra/cqlshrc
type Bookmark struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	// Keyspace is the default keyspace
	Keyspace string `json:"keyspace"`
	// Keyspaces is the list of common keyspaces
	Keyspaces []string        `json:"keyspaces"`
	Templates []QueryTemplate `json:"templates"`
}

type QueryType string

const (
	QueryFetchOneRow  QueryType = "fetchOneRow"
	QueryUpdateOneRow QueryType = "updateOneRow"
	// QueryDDL is for DDL like create/alter/drop table/keyspace etc.
	QueryDDL QueryType = "ddl"
)

// QueryTemplate contains a query and its parameters definitions
type QueryTemplate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// Query is the query string with place holders
	Query  string           `json:"query"`
	Params []QueryParameter `json:"params"`
	Type   QueryType        `json:"type"`
}

type ParamType string

const (
	ParamInt    ParamType = "int"
	ParamString ParamType = "string"
)

// QueryParameter is loosely modeled after swagger's schema definition
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#schema-object
type QueryParameter struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        ParamType `json:"type"`
	Required    bool      `json:"required"`
}
