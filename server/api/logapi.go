package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tsaikd/full-stack-demo/server/slacklog"
)

type logType struct {
}

type logContext struct {
	param interface{}
}

// logapiHandler log API to elastic
func logapiHandler(logTypeName LogTypeName, r Render, start time.Time, next func(logctx *logContext)) {
	logctx := &logContext{}

	defer func(logctx *logContext) {
		panic := ""
		if err := recover(); err != nil {
			if !r.IsWritten() {
				r.WriteResponse(nil, http.StatusInternalServerError, "500 Internal Server Error")
			}
			slacklog.LogAsync(fmt.Sprintf("%s\n%s", err, stack(3)))
		}
		logapi(logTypeName, r, logctx.param, time.Now().Sub(start), panic)
	}(logctx)

	next(logctx)
}

func logapi(
	logTypeName LogTypeName,
	r Render,
	param interface{},
	duration time.Duration,
	panic string,
) {
}
