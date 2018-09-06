package applog

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/server/appconst"
)

// command line flags
var (
	flagDebug = &cobrather.BoolFlag{
		Name:    "debug",
		Default: false,
		Usage:   "Show debug level message",
		EnvVar:  appconst.EnvPrefix + "DEBUG",
	}
)

// Module info
var Module = &cobrather.Module{
	Use: "applog",
	GlobalFlags: []cobrather.Flag{
		flagDebug,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		gDebug = flagDebug.Bool()
		return nil
	},
}
