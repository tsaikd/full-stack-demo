package elastic

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/server/appconst"
)

// command line flags
var (
	flagElasticPrefix = &cobrather.StringFlag{
		Name:    "elastic",
		Default: "",
		Usage:   "Elastic web service url prefix, e.g. http://localhost:9200",
		EnvVar:  appconst.EnvPrefix + "ELASTIC",
	}
)

// Module info
var Module = &cobrather.Module{
	Use: "elastic",
	Flags: []cobrather.Flag{
		flagElasticPrefix,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return Config(ctx, flagElasticPrefix.String())
	},
	PostRunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return Close()
	},
}
