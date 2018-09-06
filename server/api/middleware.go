package api

import (
	"io"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tsaikd/KDGoLib/webutil/render/httpRender"
	"github.com/tsaikd/full-stack-demo/server/slacklog"
)

// Render wrap http render interface
type Render interface {
	httpRender.Render
}

func newRender(w http.ResponseWriter, req *http.Request) Render {
	return httpRender.New(w, req,
		httpRender.OptionErrorPathTrimPrefix("github.com/tsaikd/full-stack-demo/"),
	)
}

// ProvideRender is a http middleware handler to provide Render and Token with Log function
func ProvideRender(w http.ResponseWriter, req *http.Request, logTypeName LogTypeName, param interface{}, next func(r Render)) {
	start := time.Now()
	r := newRender(w, req)

	logapiHandler(logTypeName, r, start, func(logctx *logContext) {
		logctx.param = param

		if logctx.param != nil {
			if err := BindParam(logctx.param, req); err != nil {
				r.Error(err)
				slacklog.Trace(err)
				return
			}
		}

		next(r)
	})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		// don't return errors to maintain backwards compatibility
		slacklog.Trace(reason)
	},
	CheckOrigin: func(r *http.Request) bool {
		// allow all connections by default
		return true
	},
}

// ProvideWebSocket is a http middleware handler to provide websocket.Conn
func ProvideWebSocket(w http.ResponseWriter, req *http.Request, next func(ws *websocket.Conn)) {
	start := time.Now()

	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		// log websocket error
		ProvideError(w, req, LogTypeNameAPI, start, err)
		return
	}
	defer func() {
		slacklog.Trace(ws.Close())
	}()

	next(ws)
}

// ProvideHandler is a http middleware handler to wrap http.Handler with Log function
func ProvideHandler(w http.ResponseWriter, req *http.Request, logTypeName LogTypeName, handler http.HandlerFunc, start time.Time) {
	r := newRender(w, req)

	logapiHandler(logTypeName, r, start, func(logctx *logContext) {
		handler(r.GetResponseWriter(), r.GetRequest())
	})
}

// ProvideContent is a http middleware handler to wrap http.ServeContent with Log function
func ProvideContent(w http.ResponseWriter, req *http.Request, logTypeName LogTypeName, name string, modtime time.Time, content io.ReadSeeker, start time.Time) {
	r := newRender(w, req)

	logapiHandler(logTypeName, r, start, func(logctx *logContext) {
		http.ServeContent(r.GetResponseWriter(), r.GetRequest(), name, modtime, content)
	})
}

// ProvideError is a http middleware handler to response error
func ProvideError(w http.ResponseWriter, req *http.Request, logTypeName LogTypeName, start time.Time, errOut error) {
	r := newRender(w, req)

	logapiHandler(logTypeName, r, start, func(logctx *logContext) {
		r.Error(errOut)
	})
}
