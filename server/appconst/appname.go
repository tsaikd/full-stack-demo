package appconst

import (
	"os"
	"strings"

	"github.com/tsaikd/KDGoLib/version"
)

// AppName application name
var AppName = "full-stack-demo"

// AppNameLower application lowercase name
var AppNameLower = strings.ToLower(AppName)

// EnvPrefix environment prefix
var EnvPrefix = "FSD_"

// Hostname of running application
var Hostname = func() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}()

// AppNameVer application name with version info and hostname
var AppNameVer = AppName + " " + version.VERSION + " @" + Hostname
