package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetVersionFromSource(t *testing.T) {
	require := require.New(t)
	require.NotNil(require)

	version, err := GetVersionFromSource()
	require.NoError(err)
	require.NotEmpty(version.Version)
	require.NotEmpty(version.GoVersion)
	require.NotEmpty(version.BuildTime)
	require.NotEmpty(version.GitCommit)
}
