# gocql

https://github.com/gocql/gocql

`````go
// schema metadata for a keyspace
type KeyspaceMetadata struct {
	Name            string
	DurableWrites   bool
	StrategyClass   string
	StrategyOptions map[string]interface{}
	Tables          map[string]*TableMetadata
}

// schema metadata for a table (a.k.a. column family)
type TableMetadata struct {
	Keyspace          string
	Name              string
	KeyValidator      string
	Comparator        string
	DefaultValidator  string
	KeyAliases        []string
	ColumnAliases     []string
	ValueAlias        string
	PartitionKey      []*ColumnMetadata
	ClusteringColumns []*ColumnMetadata
	Columns           map[string]*ColumnMetadata
	OrderedColumns    []string
}

// schema metadata for a column
type ColumnMetadata struct {
	Keyspace        string
	Table           string
	Name            string
	ComponentIndex  int
	Kind            ColumnKind
	Validator       string
	Type            TypeInfo
	ClusteringOrder string
	Order           ColumnOrder
	Index           ColumnIndexMetadata
}
`````