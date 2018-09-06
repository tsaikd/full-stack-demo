package coinHistory_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsaikd/full-stack-demo/server/database/coinHistory"
	"github.com/tsaikd/full-stack-demo/server/database/coinHistory/coinHistoryTest"
	"github.com/tsaikd/full-stack-demo/server/database/dbTest"
)

func Test_coins(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	ctx := context.Background()
	testCtx := dbTest.NewTestContext(ctx, t)
	defer testCtx.Close()
	tx := testCtx.Tx
	require.NotNil(tx)

	coinHistory1 := coinHistoryTest.New(testCtx, coinHistoryTest.CoinHistory{
		Symbol:    "mockBTC",
		Timestamp: time.Now(),
		Price:     3.14,
	})
	defer coinHistory1.Close()

	if result, err := coinHistory.Get(ctx, tx, "mockBTC"); assert.NoError(err) {
		require.Equal("mockBTC", result[0].Symbol)
		require.NotEmpty(result[0].Timestamp)
		require.EqualValues(3.14, result[0].Price)
	}
}
