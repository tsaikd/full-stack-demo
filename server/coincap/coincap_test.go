package coincap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHistory1Day(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	if result, err := History1Day("BTC"); assert.NoError(err) {
		require.NotEmpty(result.Price)
	}
}
