# gocql

https://github.com/gocql/gocql

````go
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
````

````go
type ColumnInfo struct {
	Keyspace string
	Table    string
	Name     string
	TypeInfo TypeInfo
}

type ColumnInfo struct {
	Keyspace string
	Table    string
	Name     string
	TypeInfo TypeInfo
}

type TypeInfo interface {
	Type() Type
	Version() byte
	Custom() string

	// New creates a pointer to an empty version of whatever type
	// is referenced by the TypeInfo receiver
	New() interface{}
}

// String returns a human readable name for the Cassandra datatype
// described by t.
// Type is the identifier of a Cassandra internal datatype.
type Type int

const (
	TypeCustom    Type = 0x0000
	TypeAscii     Type = 0x0001
	TypeBigInt    Type = 0x0002
	TypeBlob      Type = 0x0003
	TypeBoolean   Type = 0x0004
	TypeCounter   Type = 0x0005
	TypeDecimal   Type = 0x0006
	TypeDouble    Type = 0x0007
	TypeFloat     Type = 0x0008
	TypeInt       Type = 0x0009
	TypeText      Type = 0x000A
	TypeTimestamp Type = 0x000B
	TypeUUID      Type = 0x000C
	TypeVarchar   Type = 0x000D
	TypeVarint    Type = 0x000E
	TypeTimeUUID  Type = 0x000F
	TypeInet      Type = 0x0010
	TypeDate      Type = 0x0011
	TypeTime      Type = 0x0012
	TypeSmallInt  Type = 0x0013
	TypeTinyInt   Type = 0x0014
	TypeDuration  Type = 0x0015
	TypeList      Type = 0x0020
	TypeMap       Type = 0x0021
	TypeSet       Type = 0x0022
	TypeUDT       Type = 0x0030
	TypeTuple     Type = 0x0031
)

type NativeType struct {
	proto  byte
	typ    Type
	custom string // only used for TypeCustom
}

type CollectionType struct {
	NativeType
	Key  TypeInfo // only used for TypeMap
	Elem TypeInfo // only used for TypeMap, TypeList and TypeSet
}

type TupleTypeInfo struct {
	NativeType
	Elems []TypeInfo
}

type UDTField struct {
	Name string
	Type TypeInfo
}

type UDTTypeInfo struct {
	NativeType
	KeySpace string
	Name     string
	Elements []UDTField
}
````