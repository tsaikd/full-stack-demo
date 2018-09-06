package apijs

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/cmd"
	"github.com/tsaikd/full-stack-demo/server/applog"
)

// Module info
var Module = &cobrather.Module{
	Use:   "apijs",
	Short: "Update all registed API to javascript library",
	Dependencies: []*cobrather.Module{
		applog.Module,
		cmd.Module,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return nil
	},
}
