package apiModule

//go:generate full-stack-demo apigen --resource "/version"

import (
	"context"

	"github.com/tsaikd/KDGoLib/version"
	"github.com/tsaikd/full-stack-demo/server/api"
)

func apiFunc(
	ctx context.Context,
	r api.Render,
	areq *apiReq,
) (err error) {
	versionInfo := version.Get()

	// remove godeps data in api response
	versionInfo.Godeps = nil

	r.JSON(apiRes{versionInfo})
	return nil
}
