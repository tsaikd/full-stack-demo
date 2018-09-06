package elastic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_index_non_config(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	ctx := context.Background()
	err := Config(ctx, "-")
	require.NoError(err)

	Elastic().Index("idxname", "idxtype", "", nil)

	err = Close()
	require.NoError(err)
}
