# Schema

How to query schema using CQL (not cqlsh), it's different in C2 and C3

- gocql has `KeyspaceMetadata` see https://github.com/gocql/gocql/blob/master/metadata.go
- [ ] TODO: list keyspace is not supported in gocql
- https://godoc.org/github.com/gocql/gocql#Session.KeyspaceMetadata
  - KeyspaceMetadata returns the schema metadata for the keyspace specified. Returns an error if the keyspace does not exist
  
## C2

https://docs.datastax.com/en/cql/3.1/cql/cql_using/use_query_system_c.html

````sql
select * from system.schema_keyspaces;
````

- generated from https://jmalarcon.github.io/markdowntables/

|keyspace_name|durable_writes|strategy_class|strategy_options|
|--- |--- |--- |--- |
|system_auth|true|org.apache.cassandra.locator.SimpleStrategy|{"replication_factor":"1"}|
|system_distributed|true|org.apache.cassandra.locator.SimpleStrategy|{"replication_factor":"3"}|
|system|true|org.apache.cassandra.locator.LocalStrategy|{}|
|system_traces|true|org.apache.cassandra.locator.SimpleStrategy|{"replication_factor":"2"}|


## C3

https://docs.datastax.com/en/cql/3.3/cql/cql_using/useQuerySystem.html

````sql
select * from system_schema.keyspaces;
````

|keyspace_name|durable_writes|replication|
|--- |--- |--- |
|system_auth|true|{"class": "org.apache.cassandra.locator.SimpleStrategy","replication_factor": "1"}|
|system_schema|true|{"class": "org.apache.cassandra.locator.LocalStrategy"}|
|system_distributed|true|{"class": "org.apache.cassandra.locator.SimpleStrategy","replication_factor": "3"}|
|system|true|{"class": "org.apache.cassandra.locator.LocalStrategy"}|
|system_traces|true|{"class": "org.apache.cassandra.locator.SimpleStrategy","replication_factor": "2"}|
