package types

import (
	"github.com/gocql/gocql"
)

type QueryRequest struct {
	Host     string `json:"host"`
	Keyspace string `json:"keyspace"`
	Query    string `json:"query"`
}

type QueryResult struct {
	Err     error // TODO: change it to another struct maybe ....
	Columns []gocql.ColumnInfo
	Rows    []map[string]interface{}
}
