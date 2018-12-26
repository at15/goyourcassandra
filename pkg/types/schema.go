package types

import "github.com/gocql/gocql"

type KeyspaceMetadata struct {
	Name            string                    `json:"name"`
	DurableWrites   bool                      `json:"durableWrites"`
	StrategyClass   string                    `json:"strategyClass"`
	StrategyOptions map[string]interface{}    `json:"strategyOptions"`
	Tables          map[string]*TableMetadata `json:"tables"`
}

type TableMetadata struct {
	Keyspace          string                     `json:"keyspace"`
	Name              string                     `json:"name"`
	KeyValidator      string                     `json:"keyValidator"`
	Comparator        string                     `json:"comparator"`
	DefaultValidator  string                     `json:"defaultValidator"`
	KeyAliases        []string                   `json:"keyAliases"`
	ColumnAliases     []string                   `json:"columnAliases"`
	ValueAlias        string                     `json:"valueAlias"`
	PartitionKey      []*ColumnMetadata          `json:"partitionKey"`
	ClusteringColumns []*ColumnMetadata          `json:"clusteringColumns"`
	Columns           map[string]*ColumnMetadata `json:"columns"`
	OrderedColumns    []string                   `json:"orderedColumns"`
}

type ColumnMetadata struct {
	Keyspace        string              `json:"keyspace"`
	Table           string              `json:"table"`
	Name            string              `json:"name"`
	ComponentIndex  int                 `json:"componentIndex"`
	Kind            string              `json:"kind"` // NOTE: was ColumnKind, partition_key, clustering_key, regular, compact, static
	Validator       string              `json:"validator"`
	Type            DataType            `json:"type"` // NOTE: was TypeInfo
	ClusteringOrder string              `json:"clusteringOrder"`
	Order           string              `json:"order"` // NOTE: was bool, asc is false, desc is true
	Index           ColumnIndexMetadata `json:"index"`
}

type ColumnIndexMetadata struct {
	Name    string                 `json:"name"`
	Type    string                 `json:"type"`
	Options map[string]interface{} `json:"options"`
}

// FIXME: hand written deep copier

func CopyKeyspaceMetadata(o *gocql.KeyspaceMetadata) KeyspaceMetadata {
	c := KeyspaceMetadata{
		Name:            o.Name,
		DurableWrites:   o.DurableWrites,
		StrategyClass:   o.StrategyClass,
		StrategyOptions: o.StrategyOptions,
	}
	c.Tables = make(map[string]*TableMetadata, len(o.Tables))
	for k, tbl := range o.Tables {
		t := CopyTableMetadata(tbl)
		c.Tables[k] = &t
	}
	return c
}

func CopyTableMetadata(o *gocql.TableMetadata) TableMetadata {
	c := TableMetadata{
		Keyspace:         o.Keyspace,
		Name:             o.Name,
		KeyValidator:     o.KeyValidator,
		Comparator:       o.Comparator,
		DefaultValidator: o.DefaultValidator,
		KeyAliases:       o.KeyAliases,
		ColumnAliases:    o.ColumnAliases,
		ValueAlias:       o.ValueAlias,
		OrderedColumns:   o.OrderedColumns,
	}
	for _, pk := range o.PartitionKey {
		t := CopyColumnMetadata(pk)
		c.PartitionKey = append(c.PartitionKey, &t)
	}
	for _, cc := range o.ClusteringColumns {
		t := CopyColumnMetadata(cc)
		c.ClusteringColumns = append(c.ClusteringColumns, &t)
	}
	c.Columns = make(map[string]*ColumnMetadata, len(o.Columns))
	for k, col := range o.Columns {
		t := CopyColumnMetadata(col)
		c.Columns[k] = &t
	}
	return c
}

func CopyColumnMetadata(o *gocql.ColumnMetadata) ColumnMetadata {
	return ColumnMetadata{
		Keyspace:        o.Keyspace,
		Table:           o.Table,
		Name:            o.Name,
		ComponentIndex:  o.ComponentIndex,
		Kind:            o.Kind.String(),
		Validator:       o.Validator,
		Type:            MustDataType(o.Type),
		ClusteringOrder: o.ClusteringOrder,
		Order:           CopyColumnOrder(o.Order),
		Index:           CopyColumnIndexMetadata(o.Index),
	}
}

func CopyColumnOrder(o gocql.ColumnOrder) string {
	switch o {
	case gocql.ASC:
		return "asc"
	case gocql.DESC:
		return "desc"
	default:
		return "unknown_order"
	}
}

func CopyColumnIndexMetadata(o gocql.ColumnIndexMetadata) ColumnIndexMetadata {
	return ColumnIndexMetadata{
		Name:    o.Name,
		Type:    o.Type,
		Options: o.Options,
	}
}
