package apiModule

import (
	"net/http"

	"github.com/tsaikd/KDGoLib/version"
	"github.com/tsaikd/full-stack-demo/server/api"
)

func init() {
	// Get version info
	api.Add(&api.API{
		Method:  api.MethodGET,
		Pattern: "/version",
		Handler: func(w http.ResponseWriter, req *http.Request) {
			areq := apiReq{}
			api.ProvideRender(w, req, api.LogTypeNameAPI, &areq, func(r api.Render) {
				ctx := r.GetRequest().Context()
				if err := apiFunc(
					ctx,
					r,
					&areq,
				); err != nil && r.GetError() == nil {
					r.Error(err)
				}
			})
		},
	})
}

type apiReq struct {
}

type apiRes struct {
	version.Version
}
