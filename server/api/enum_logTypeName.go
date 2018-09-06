package api

import (
	"database/sql/driver"

	"github.com/tsaikd/KDGoLib/enumutil"
)

// LogTypeName main type
type LogTypeName int8

// List all valid enum
const (
	LogTypeNameAPI LogTypeName = 1 + iota
	LogTypeNamePage
)

var factoryLogTypeName = enumutil.NewEnumFactory().
	Add(LogTypeNameAPI, "api").
	Add(LogTypeNamePage, "page").
	Build()

func (t LogTypeName) String() string {
	return factoryLogTypeName.String(t)
}

// MarshalJSON return jsonfy []byte of enum
func (t LogTypeName) MarshalJSON() ([]byte, error) {
	return factoryLogTypeName.MarshalJSON(t)
}

// UnmarshalJSON decode json data to enum
func (t *LogTypeName) UnmarshalJSON(b []byte) (err error) {
	return factoryLogTypeName.UnmarshalJSON(t, b)
}

// Scan decode SQL value to enum
func (t *LogTypeName) Scan(value interface{}) (err error) {
	return factoryLogTypeName.Scan(t, value)
}

// Value return enum data for SQL
func (t LogTypeName) Value() (v driver.Value, err error) {
	return factoryLogTypeName.Value(t)
}
