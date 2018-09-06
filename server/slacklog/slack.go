package slacklog

import (
	"github.com/nlopes/slack"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/full-stack-demo/server/appconst"
	"github.com/tsaikd/full-stack-demo/server/applog"
)

// errors
var (
	ErrSlackAuthTestFailed = errutil.NewFactory("slack client auth test failed")
	ErrSkipSlackLog        = errutil.NewFactory("skip slack log, use console log instead")
)

var (
	slackclient        *slack.Client
	slackNotifyChannel string

	// used for notify
	logBotName     = appconst.Hostname + "/" + appconst.AppName
	traceFormatter = &errutil.ConsoleFormatter{
		Seperator:  "; ",
		TimeFormat: "2006-01-02 15:04:05 ",
		LongFile:   true,
		Line:       true,
	}
)

// Config init slack client
func Config(token string, notifyChannel string) (err error) {
	switch token {
	case "", "-":
		applog.Logger().Debug(ErrSkipSlackLog.New(nil).Error())
		return nil
	}

	slackclient = slack.New(token)
	if _, err = slackclient.AuthTest(); err != nil {
		return ErrSlackAuthTestFailed.New(err)
	}

	slackNotifyChannel = notifyChannel

	return
}

// LogAsync run Log async
func LogAsync(msg string) {
	go func() {
		applog.Trace(Log(msg))
	}()
}

// Log send message to slack log channel
func Log(msg string) (err error) {
	if slackclient == nil {
		return nil
	}
	param := slack.NewPostMessageParameters()
	param.Username = logBotName
	_, _, err = slackclient.PostMessage(slackNotifyChannel, msg, param)
	return
}

// Trace error stack, output to slack and fallback to console
func Trace(err error) {
	if err == nil {
		return
	}

	if slackclient == nil {
		applog.Trace(err)
		return
	}

	msg, errFmt := traceFormatter.FormatSkip(err, 1)
	if errFmt != nil {
		applog.Trace(errFmt)
		return
	}
	LogAsync(msg)
}
