package types

import "github.com/gocql/gocql"

// QueryRequest can be send using POST request if there are sensitive information in query
// Or just use a GET request if you shared it in a trusted environment
type QueryRequest struct {
	Host     string `json:"host"`
	Keyspace string `json:"keyspace"`
	Query    string `json:"query"`
	// Bindings is used for prepared query
	Bindings []interface{} `json:"bindings"`
}

type QueryResult struct {
	Err     error                    // TODO: change it to another struct, or add a error field
	Columns []Column                 `json:"columns"`
	Rows    []map[string]interface{} `json:"rows"`
}

// Column replace gocql.TypeInfo with DataType
type Column struct {
	Keyspace string   `json:"keyspace"`
	Table    string   `json:"table"`
	Name     string   `json:"name"`
	Type     DataType `json:"type"`
}

func CopyColumns(cols []gocql.ColumnInfo) []Column {
	copied := make([]Column, len(cols))
	for i, col := range cols {
		copied[i] = Column{
			Keyspace: col.Keyspace,
			Table:    col.Table,
			Name:     col.Name,
			Type:     MustDataType(col.TypeInfo),
		}
	}
	return copied
}
