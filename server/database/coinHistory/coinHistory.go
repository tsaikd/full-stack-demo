package coinHistory

import (
	"context"
	"time"

	"github.com/tsaikd/KDGoLib/sqlutil"
	"github.com/tsaikd/KDGoLib/sqlutil/dbtypes"
	"github.com/tsaikd/full-stack-demo/server/ramltype"
)

// TableName contains table info
var TableName = struct {
	CoinHistory string
}{
	CoinHistory: "coinHistory",
}

// Insert into database table
func Insert(
	ctx context.Context,
	tx dbtypes.Named,
	symbol string,
	timestamp time.Time,
	price float64,
) (coinHistoryID ramltype.UUID, err error) {
	err = sqlutil.NamedGetContext(ctx, tx, &coinHistoryID, `
INSERT INTO coinHistory
	(symbol, timestamp, price)
VALUES
	(:symbol, :timestamp, :price)
RETURNING coinHistoryID
	;`, map[string]interface{}{
		"symbol":    symbol,
		"timestamp": timestamp,
		"price":     price,
	})
	return
}

// Delete coin from database
func Delete(
	ctx context.Context,
	tx dbtypes.Named,
	coinHistoryID ramltype.UUID,
) (err error) {
	_, err = sqlutil.NamedExecStrictContext(ctx, tx, `
DELETE FROM coinHistory
WHERE coinHistoryID = :coinHistoryID
	;`, map[string]interface{}{
		"coinHistoryID": coinHistoryID,
	})
	return
}

// Get from database table
func Get(
	ctx context.Context,
	tx dbtypes.Named,
	symbol string,
) (result []ramltype.ResponseCoinHistoryItem, err error) {
	err = sqlutil.NamedSelectContext(ctx, tx, &result, `
SELECT
	CoinHistoryID,
	Symbol,
	Timestamp,
	Price
FROM coinHistory
WHERE symbol = :symbol
ORDER BY timestamp ASC
	;`, map[string]interface{}{
		"symbol": symbol,
	})
	return
}
