package apiModule

import (
	"net/http"

	"github.com/tsaikd/KDGoLib/sqlutil/dbtypes"
	"github.com/tsaikd/full-stack-demo/server/api"
	"github.com/tsaikd/full-stack-demo/server/database"
	"github.com/tsaikd/full-stack-demo/server/ramltype"
)

func init() {
	// List fund
	api.Add(&api.API{
		Method:  api.MethodGET,
		Pattern: "/get/coinHistory",
		Handler: func(w http.ResponseWriter, req *http.Request) {
			areq := apiReq{}
			api.ProvideRender(w, req, api.LogTypeNameAPI, &areq, func(r api.Render) {
				ctx := r.GetRequest().Context()
				database.ProvideTransaction(r, func(tx dbtypes.Transactionx) (rollback bool) {
					if err := apiFunc(
						ctx,
						r,
						&areq,
						tx,
					); err != nil && r.GetError() == nil {
						r.Error(err)
					}
					return r.GetError() != nil
				})
			})
		},
	})
}

type apiReq struct {
	Symbol string `json:"symbol" valid:"required"`
}

type apiRes = ramltype.ResponseCoinHistory
