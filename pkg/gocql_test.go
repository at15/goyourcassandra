package pkg

import (
	"fmt"
	"testing"

	"github.com/gocql/gocql"
)

// https://github.com/gocql/gocql
// TODO: gocql does not allow using use statement to switch keyspace
// - would desc keyspace work
// - for different keyspace, need different connection

// https://github.com/gocql/gocql#important-default-keyspace-changes
// gocql no longer supports executing "use " statements to simplify the library.
// The user still has the ability to define the default keyspace for connections
// but now the keyspace can only be defined before a session is created.
// Queries can still access keyspaces by indicating the keyspace in the query:

func TestListKeyspace(t *testing.T) {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	if err != nil {
		t.Fatal(err)
	}
	// FIXME: desc is a cqlsh specific command ... it's queried from system keyspace ...
	// https://docs.datastax.com/en/cql/3.1/cql/cql_using/use_query_system_tables_t.html
	//descKeyspace := "desc keyspaces;"
	descKeyspace := "SELECT * from system.schema_keyspaces;"
	if err := session.Query(descKeyspace).Exec(); err != nil {
		t.Fatal(err)
	}
	iter := session.Query(descKeyspace).Iter()
	fmt.Println(iter.Columns())
	for {
		// New map each iteration
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			break
		}
		// Do things with row
		if ks, ok := row["keyspace_name"]; ok {
			fmt.Printf("keyspace_name: %v\n", ks)
		}
	}
}
