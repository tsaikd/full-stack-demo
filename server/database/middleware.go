package database

import (
	"github.com/tsaikd/KDGoLib/sqlutil/dbtypes"
	"github.com/tsaikd/KDGoLib/webutil/render"
	"github.com/tsaikd/full-stack-demo/server/slacklog"
)

// ProvideTransaction is a http middleware handler to provide tx
func ProvideTransaction(r render.Error, next func(tx dbtypes.Transactionx) (rollback bool)) {
	tx, err := db.Beginx()
	if err != nil {
		r.Error(err)
		return
	}

	if next(tx) || r.GetError() != nil {
		slacklog.Trace(tx.Rollback())
	} else {
		if err = tx.Commit(); err != nil {
			r.Error(err)
		}
	}
}
