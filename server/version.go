package server

import (
	"github.com/tsaikd/KDGoLib/version"
	"github.com/tsaikd/full-stack-demo/server/appconst"
	"github.com/tsaikd/full-stack-demo/server/util"
)

func init() {
	version.NAME = appconst.AppName

	if version.GITCOMMIT == "" {
		ver, err := util.GetVersionFromSource()
		if err != nil {
			panic(err)
		}

		version.VERSION = ver.Version
		version.GITCOMMIT = ver.GitCommit
	}
}
