package database

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/server/appconst"
)

// command line flags
var (
	flagDBConnectionURL = &cobrather.StringFlag{
		Name:    "dburl",
		Default: "mysql://USER:PASS@tcp(HOST:PORT)/DBNAME",
		Usage:   "Database connection URL",
		EnvVar:  appconst.EnvPrefix + "DBCONNURL",
	}
)

// Module info
var Module = &cobrather.Module{
	Use: "database",
	Flags: []cobrather.Flag{
		flagDBConnectionURL,
	},
	RunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		connurl := flagDBConnectionURL.String()
		ssldir := appconst.ProjectRootDir
		return Config(connurl, ssldir)
	},
	PostRunE: func(ctx context.Context, cmd *cobra.Command, args []string) error {
		return Close()
	},
}

// DBConnectionURL return cli arg value of flagDBConnectionURL
func DBConnectionURL() string {
	return flagDBConnectionURL.String()
}
