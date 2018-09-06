package main

import (
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/cmd/apigen"
	"github.com/tsaikd/full-stack-demo/cmd/apijs"
	"github.com/tsaikd/full-stack-demo/cmd/fetchCoin"
	"github.com/tsaikd/full-stack-demo/cmd/modimport"
	"github.com/tsaikd/full-stack-demo/server"
	"github.com/tsaikd/full-stack-demo/server/appconst"
)

// Module info
var Module = &cobrather.Module{
	Use: appconst.AppName,
	Commands: []*cobrather.Module{
		apigen.Module,
		apijs.Module,
		fetchCoin.Module,
		modimport.Module,
		server.Module,
		cobrather.VersionModule,
	},
}

func main() {
	Module.MustMainRun()
}
