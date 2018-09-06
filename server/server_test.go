package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsaikd/full-stack-demo/server/api"
	"github.com/tsaikd/full-stack-demo/server/webclient"
)

func Test_newRouter(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(assert)
	require := require.New(t)
	require.NotNil(require)

	ts := httptest.NewServer(newRouter())
	defer ts.Close()

	if body, res, err := webclient.GetString(ts.URL + "/api/1/version"); assert.NoError(err) {
		require.EqualValues(http.StatusOK, res.StatusCode)
		require.Contains(body, "version")
		require.Contains(body, "buildtime")
		require.Contains(body, "gitcommit")
	}

	if _, res, err := webclient.GetString(ts.URL + "/api/2/version"); assert.Error(err) {
		require.EqualValues(http.StatusNotFound, res.StatusCode)
		require.Contains(err.Error(), api.ErrAPINotFound.Error())
	}
}
