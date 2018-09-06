package slacklog

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/server/appconst"
)

// command line flags
var (
	flagSlackAPIToken = &cobrather.StringFlag{
		Name:   "slackapitoken",
		Usage:  "Slack API Token",
		EnvVar: appconst.EnvPrefix + "SLACKAPI_TOKEN",
	}
	flagSlackNotifyChannel = &cobrather.StringFlag{
		Name:    "slacknotifychannel",
		Default: "#general",
		Usage:   "Slack notify channel",
		EnvVar:  appconst.EnvPrefix + "SLACKAPI_NOTIFY_CHANNEL",
	}
)

// Module info
var Module = &cobrather.Module{
	Use: "slacklog",
	Flags: []cobrather.Flag{
		flagSlackAPIToken,
		flagSlackNotifyChannel,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return Config(flagSlackAPIToken.String(), flagSlackNotifyChannel.String())
	},
}
