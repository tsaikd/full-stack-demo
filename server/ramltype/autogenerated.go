package ramltype

import (
	"database/sql/driver"
	"time"

	"github.com/tsaikd/KDGoLib/sqlutil"
)

// ResponseCoinHistoryItem autogenerated raml type
type ResponseCoinHistoryItem struct {
	CoinHistoryID string    `json:"coinHistoryID"`
	Symbol        string    `json:"symbol"`
	Timestamp     time.Time `json:"timestamp"`
	Price         float64   `json:"price"`
}

// ResponseCoinHistory autogenerated raml type
type ResponseCoinHistory = []ResponseCoinHistoryItem

// UUID autogenerated raml type
type UUID string

// Scan decode SQL value to enum
func (t *UUID) Scan(value interface{}) (err error) {
	return sqlutil.SQLScanString(t, value)
}

// Value return enum data for SQL
func (t UUID) Value() (v driver.Value, err error) {
	if t == "" {
		return nil, nil
	}
	return string(t), nil
}