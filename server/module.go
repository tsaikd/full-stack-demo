package server

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/server/api"
	"github.com/tsaikd/full-stack-demo/server/appconst"
	"github.com/tsaikd/full-stack-demo/server/applog"
	"github.com/tsaikd/full-stack-demo/server/database"
	"github.com/tsaikd/full-stack-demo/server/slacklog"
)

// command line flags
var (
	flagAddress = &cobrather.StringFlag{
		Name:    "address",
		Default: ":3000",
		Usage:   "Server listen host:port",
		EnvVar:  appconst.EnvPrefix + "ADDRESS",
	}
	flagTLS = &cobrather.BoolFlag{
		Name:   "tls",
		Usage:  "API with TLS, 'cert.pem', 'key.pem' should be prepared",
		EnvVar: appconst.EnvPrefix + "TLS",
	}
)

// Module info
var Module = &cobrather.Module{
	Use:   "server",
	Short: "web server",
	Dependencies: []*cobrather.Module{
		api.Module,
		applog.Module,
		database.Module,
		slacklog.Module,
	},
	Flags: []cobrather.Flag{
		flagAddress,
		flagTLS,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		webIndexContent, webIndexMetaContent, webIndexModified = initWebIndex()
		webJSContent, webJSModified = initWebJS()

		return run(ctx, flagAddress.String(), flagTLS.Bool())
	},
}
