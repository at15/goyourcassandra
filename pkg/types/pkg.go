package types

import (
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/gocql/gocql"
)

var log = dlog.NewLibraryLogger()

// DataType is a merged struct of gocql.TypeInfo implementations
type DataType struct {
	// Type is for native simple type like int
	Type string `json:"type"`
	// Key is only used for map
	Key string `json:"key"`
	// Elem is used for map, list and set
	Elem string `json:"elem"`
	// Elems is used for tuple
	Elems []string `json:"elems"`
}

// TODO: we didn't handle nested frozen collection, i.e. map<int, <frozen map<int, string>> which can't have update to individual fields
// https://docs.datastax.com/en/cql/3.3/cql/cql_reference/collection_type_r.html
func ToDataType(t gocql.TypeInfo) (DataType, error) {
	dt := DataType{Type: t.Type().String()}
	switch t.(type) {
	case gocql.NativeType:
		return dt, nil
	case gocql.CollectionType:
		ct := t.(gocql.CollectionType)
		switch ct.Type() {
		case gocql.TypeMap:
			dt.Key = ct.Key.Type().String()
			dt.Elem = ct.Elem.Type().String()
		case gocql.TypeList, gocql.TypeSet:
			dt.Elem = ct.Elem.Type().String()
		}
		return dt, nil
	case gocql.TupleTypeInfo:
		tt := t.(gocql.TupleTypeInfo)
		elemes := tt.Elems
		for _, eleme := range elemes {
			dt.Elems = append(dt.Elems, eleme.Type().String())
		}
	}
	return DataType{Type: t.Type().String()}, errors.Errorf("unsupported type %s custom %s", t.Type().String(), t.Custom())
}

// log and ignore the error
func MustDataType(t gocql.TypeInfo) DataType {
	dt, err := ToDataType(t)
	if err != nil {
		log.Warnf("error convert gocql.TypeInfo to DataType: %s", err)
	}
	return dt
}
