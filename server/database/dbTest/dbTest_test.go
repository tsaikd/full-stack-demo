package dbTest_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsaikd/full-stack-demo/server/database/dbTest"
)

func TestNewTestContext(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	ctx := context.Background()
	testCtx := dbTest.NewTestContext(ctx, t)
	defer testCtx.Close()
	require.NotNil(testCtx)
	require.NotNil(testCtx.DB)
	require.NotNil(testCtx.Tx)
}

func TestCoalesceString(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	if str := dbTest.CoalesceString("a"); assert.NotEmpty(str) {
		require.Equal("a", str)
	}

	if str := dbTest.CoalesceString("", "b"); assert.NotEmpty(str) {
		require.Equal("b", str)
	}

	if str := dbTest.CoalesceString("", "c", ""); assert.NotEmpty(str) {
		require.Equal("c", str)
	}
}

func TestCoalesceInt64(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	if i := dbTest.CoalesceInt64(9527); assert.NotZero(i) {
		require.EqualValues(9527, i)
	}

	if i := dbTest.CoalesceInt64(0, 9527); assert.NotZero(i) {
		require.EqualValues(9527, i)
	}

	if i := dbTest.CoalesceInt64(0, 9527, 0); assert.NotZero(i) {
		require.EqualValues(9527, i)
	}
}

func TestNext(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	serial := int64(1)

	if str := dbTest.Next(&serial); assert.NotEmpty(str) {
		require.Equal("1", str)
	}

	if str := dbTest.Next(&serial); assert.NotEmpty(str) {
		require.Equal("2", str)
	}
}
