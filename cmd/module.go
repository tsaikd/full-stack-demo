package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/server/applog"
)

// Module info
var Module = &cobrather.Module{
	Use: "cmd",
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return applog.GoCmdMode()
	},
}
