package slacklog

import (
	"fmt"

	"github.com/tsaikd/full-stack-demo/server/applog"
)

// LoggerType general type for log with slack
type LoggerType interface {
	Printf(format string, v ...interface{})
}

var logger = loggerType{}

type loggerType struct{}

func (t loggerType) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if slackclient == nil {
		applog.Logger().Error(msg)
		return
	}
	applog.Trace(Log(msg))
}

// Logger return logger instance
func Logger() LoggerType {
	return logger
}
