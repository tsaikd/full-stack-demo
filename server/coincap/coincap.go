package coincap

import (
	"encoding/json"
	"time"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/full-stack-demo/server/webclient"
)

// errors
var (
	ErrInvalidFormat1 = errutil.NewFactory("invalid format: %q")
)

// ResponseMapItem response item of coincap map API
type ResponseMapItem struct {
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

// FetchMap call coincap map API
func FetchMap() (result []ResponseMapItem, err error) {
	_, err = webclient.GetJSON("http://coincap.io/map", &result)
	return
}

// ResponseHistory response of coincap history API
type ResponseHistory struct {
	Price []ResponseHistoryPrice `json:"price"`
}

// ResponseHistoryPrice price type in ResponseHistory
type ResponseHistoryPrice struct {
	Timestamp time.Time `json:"timestamp"`
	Price     float64   `json:"price"`
}

// UnmarshalJSON implement JSON unmarshaler interface
func (t *ResponseHistoryPrice) UnmarshalJSON(b []byte) (err error) {
	tmp := []interface{}{}
	if err = json.Unmarshal(b, &tmp); err != nil {
		return
	}
	if len(tmp) < 2 {
		return ErrInvalidFormat1.New(nil, string(b))
	}
	// 這裡應該要多檢查幾種 type
	ts, ok := tmp[0].(float64)
	if !ok {
		return ErrInvalidFormat1.New(nil, string(b))
	}
	t.Timestamp = time.Unix(int64(ts)/1000, 0)

	t.Price, ok = tmp[1].(float64)
	if !ok {
		return ErrInvalidFormat1.New(nil, string(b))
	}

	return
}

// History1Day call coincap /history/1day/:coin API
func History1Day(coinSymbol string) (result ResponseHistory, err error) {
	_, err = webclient.GetJSON("http://coincap.io/history/1day/"+coinSymbol, &result)
	return
}
