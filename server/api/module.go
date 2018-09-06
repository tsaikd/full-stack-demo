package api

import (
	"github.com/tsaikd/KDGoLib/cliutil/cobrather"
	"github.com/tsaikd/full-stack-demo/server/elastic"
)

// Module info
var Module = &cobrather.Module{
	Use: "api",
	Dependencies: []*cobrather.Module{
		elastic.Module,
	},
}
