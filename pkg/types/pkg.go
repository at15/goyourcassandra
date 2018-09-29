package types

import (
	"github.com/gocql/gocql"
)

type Request struct {
	Host     string `json:"host"`
	Keyspace string `json:"keyspace"`
	Query    string `json:"query"`
}

type Result struct {
	Err     error
	Columns []gocql.ColumnInfo
	Rows    []map[string]interface{}
}
