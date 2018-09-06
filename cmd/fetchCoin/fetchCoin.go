package fetchCoin

import (
	"context"
	"time"

	"github.com/tsaikd/KDGoLib/sqlutil/dbtypes"
	"github.com/tsaikd/full-stack-demo/server/applog"
	"github.com/tsaikd/full-stack-demo/server/coincap"
	"github.com/tsaikd/full-stack-demo/server/database"
	"github.com/tsaikd/full-stack-demo/server/database/coinHistory"
	pb "gopkg.in/cheggaaa/pb.v1"
)

func run(ctx context.Context, once bool) (err error) {
	if err = fetchCoin(ctx); err != nil {
		return
	}
	if once {
		return nil
	}

	// crawl per hour
	ticker := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
		if err = fetchCoin(ctx); err != nil {
			return
		}
	}
}

func fetchCoin(ctx context.Context) (err error) {
	logger := applog.Logger()
	logger.Info("Fetching data from coincap ...")
	symbol := "BTC" // symbol should be set from caller or in a loop
	result, err := coincap.History1Day("BTC")
	if err != nil {
		return
	}

	logger.Info("Saving data into database ...")
	r := &applog.ConsoleRender{}
	database.ProvideTransaction(r, func(tx dbtypes.Transactionx) (rollback bool) {
		bar := pb.New64(int64(len(result.Price))).Start()
		defer bar.Finish()

		for _, price := range result.Price {
			if _, err = coinHistory.Insert(ctx, tx, symbol, price.Timestamp, price.Price); err != nil {
				r.Error(err)
				return
			}
			bar.Increment()
		}

		return r.GetError() != nil
	})

	return
}
