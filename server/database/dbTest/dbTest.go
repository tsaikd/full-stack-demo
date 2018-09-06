package dbTest

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsaikd/full-stack-demo/server/appconst"
	"github.com/tsaikd/full-stack-demo/server/database"
)

// TestContext context for testing
type TestContext struct {
	Ctx context.Context
	T   *testing.T
	DB  *sqlx.DB

	Assert  *assert.Assertions
	Require *require.Assertions
	Tx      *sqlx.Tx
}

// NewTestContext return new TestContext instance,
// database connection info should be environment EnvPrefix_DBCONNURL
func NewTestContext(ctx context.Context, t *testing.T) *TestContext {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	connURL := os.Getenv(appconst.EnvPrefix + "DBCONNURL")
	if connURL == "" {
		t.Skip("skip DB testing")
	}

	projroot := appconst.ProjectRootDir
	require.NoError(database.Config(connURL, projroot))

	db := database.Database()
	require.NotNil(db)

	tx, err := db.BeginTxx(ctx, nil)
	require.NoError(err)
	require.NotNil(tx)

	return &TestContext{
		Ctx: ctx,
		T:   t,
		DB:  db,

		Assert:  assert,
		Require: require,
		Tx:      tx,
	}
}

// Close for closing DB connection
func (t TestContext) Close() {
	if t.T.Failed() {
		t.Require.NoError(t.Tx.Rollback())
	} else {
		t.Require.NoError(t.Tx.Commit())
	}
	t.Require.NoError(database.Close())
}

// CoalesceString return first non-empty data
func CoalesceString(list ...string) string {
	for _, str := range list {
		if str != "" {
			return str
		}
	}
	return ""
}

// CoalesceInt64 return first non-zero data
func CoalesceInt64(list ...int64) int64 {
	for _, i := range list {
		if i != 0 {
			return i
		}
	}
	return 0
}

// Next return next serial string
func Next(serial *int64) string {
	str := strconv.FormatInt(*serial, 10)
	(*serial)++
	return str
}
