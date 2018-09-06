package apiModule

//go:generate full-stack-demo apigen --resource "/get/coinHistory"

import (
	"context"

	"github.com/tsaikd/KDGoLib/sqlutil/dbtypes"
	"github.com/tsaikd/full-stack-demo/server/api"
	"github.com/tsaikd/full-stack-demo/server/database/coinHistory"
)

func apiFunc(
	ctx context.Context,
	r api.Render,
	areq *apiReq,
	tx dbtypes.Transactionx,
) (err error) {
	result, err := coinHistory.Get(ctx, tx, areq.Symbol)
	if err != nil {
		return api.ErrDatabaseGet1.New(err, coinHistory.TableName.CoinHistory)
	}

	r.JSON(apiRes(result))
	return nil
}
