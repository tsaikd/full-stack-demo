package fetchCoin

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/cmd"
	"github.com/tsaikd/full-stack-demo/server/database"
)

// command line flags
var (
	flagOnce = &cobrather.BoolFlag{
		Name:  "once",
		Usage: "Fetch only once",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:   "fetchCoin",
	Short: "Fetch coin info and insert into database",
	Dependencies: []*cobrather.Module{
		cmd.Module,
		database.Module,
	},
	Flags: []cobrather.Flag{
		flagOnce,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return run(ctx, flagOnce.Bool())
	},
}
