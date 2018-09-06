package coinHistoryTest

import (
	"time"

	"github.com/tsaikd/full-stack-demo/server/database/coinHistory"
	"github.com/tsaikd/full-stack-demo/server/database/dbTest"
	"github.com/tsaikd/full-stack-demo/server/ramltype"
)

var serial = int64(1)

// CoinHistory testing data
type CoinHistory struct {
	*dbTest.TestContext
	CoinHistoryID ramltype.UUID
	Symbol        string
	Timestamp     time.Time
	Price         float64
}

// Close for removing testing data
func (t CoinHistory) Close() {
	err := coinHistory.Delete(t.Ctx, t.Tx, t.CoinHistoryID)
	t.Require.NoError(err)
}

// New return new testing data instance
func New(testCtx *dbTest.TestContext, config CoinHistory) *CoinHistory {
	coinHistoryID, err := coinHistory.Insert(testCtx.Ctx, testCtx.Tx, config.Symbol, config.Timestamp, config.Price)
	testCtx.Require.NoError(err)

	return &CoinHistory{
		TestContext:   testCtx,
		CoinHistoryID: coinHistoryID,
		Symbol:        config.Symbol,
		Timestamp:     config.Timestamp,
		Price:         config.Price,
	}
}
