package api

import (
	"database/sql/driver"

	"github.com/tsaikd/KDGoLib/enumutil"
)

// Method main type
type Method int8

// List all valid enum
const (
	MethodGET Method = 1 + iota
	MethodPOST
)

var factoryMethod = enumutil.NewEnumFactory().
	Add(MethodGET, "GET").
	Add(MethodPOST, "POST").
	Build()

func (t Method) String() string {
	return factoryMethod.String(t)
}

// MarshalJSON return jsonfy []byte of enum
func (t Method) MarshalJSON() ([]byte, error) {
	return factoryMethod.MarshalJSON(t)
}

// UnmarshalJSON decode json data to enum
func (t *Method) UnmarshalJSON(b []byte) (err error) {
	return factoryMethod.UnmarshalJSON(t, b)
}

// Scan decode SQL value to enum
func (t *Method) Scan(value interface{}) (err error) {
	return factoryMethod.Scan(t, value)
}

// Value return enum data for SQL
func (t Method) Value() (v driver.Value, err error) {
	return factoryMethod.Value(t)
}
