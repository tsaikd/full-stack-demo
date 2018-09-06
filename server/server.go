package server

//go:generate full-stack-demo modimport --package server --dir api
//go:generate full-stack-demo apijs

import (
	"context"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	rpprof "runtime/pprof"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tsaikd/KDGoLib/version"
	"github.com/tsaikd/KDGoLib/webutil/render/httpRender"
	"github.com/tsaikd/full-stack-demo/server/api"
	"github.com/tsaikd/full-stack-demo/server/appconst"
	"github.com/tsaikd/full-stack-demo/server/applog"
	"github.com/tsaikd/full-stack-demo/server/slacklog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

type signalLogger interface {
	Info(msg string, fields ...zapcore.Field)
}

func contextWithOSSignal(parent context.Context, logger signalLogger, sig ...os.Signal) context.Context {
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, sig...)

	ctx, cancel := context.WithCancel(parent)

	go func(cancel context.CancelFunc) {
		select {
		case sig := <-osSignalChan:
			logger.Info(sig.String())
			cancel()
		}
	}(cancel)

	return ctx
}

func newRouter() *httprouter.Router {
	router := httprouter.New()

	for _, a := range api.All() {
		router.HandlerFunc(a.Method.String(), "/api/1"+a.Pattern, a.Handler)
	}

	router.HandlerFunc("GET", "/debug/pprof/", pprof.Index)
	router.HandlerFunc("GET", "/debug/pprof/cmdline", pprof.Cmdline)
	router.HandlerFunc("GET", "/debug/pprof/profile", pprof.Profile)
	router.HandlerFunc("GET", "/debug/pprof/symbol", pprof.Symbol)
	router.HandlerFunc("GET", "/debug/pprof/trace", pprof.Trace)
	for _, profile := range rpprof.Profiles() {
		router.HandlerFunc("GET", "/debug/pprof/"+profile.Name(), pprof.Index)
	}

	// handle app js to inject GCS config
	router.HandlerFunc("GET", "/static/js/app.:suffix", func(w http.ResponseWriter, req *http.Request) {
		if !strings.HasSuffix(req.RequestURI, ".js") {
			webFileServer.ServeHTTP(w, req)
			return
		}

		w.Header().Set(httpRender.HeaderCacheControl, cacheHeaderOneYear)
		http.ServeContent(w, req, filepath.Base(req.RequestURI), webJSModified, webJSContent)
	})

	router.NotFound = notFoundHandler{}

	return router
}

func newServer(address string) (server *http.Server) {
	return &http.Server{
		Addr:         address,
		Handler:      newRouter(),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
}

func listen(
	server *http.Server,
	tls bool,
) (err error) {
	logger := applog.Logger()
	slacklog.LogAsync("starting " + version.String())

	if tls {
		logger.Info(appconst.AppName+" https server start", zap.String("address", server.Addr))
		if err = server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
			switch err.(type) {
			case *os.PathError:
				genpath := filepath.Join(runtime.GOROOT(), "src", "crypto", "tls", "generate_cert.go")
				errmsg := strings.Join([]string{
					"'cert.pem' and 'key.pem' should be prepared when starting server with tls flag.",
					"To generate a development cert and key, run the following from your *nix terminal:",
					`go run ` + genpath + ` --host="localhost"`,
				}, "\n")
				logger.Error(errmsg)
				return
			}
			switch err {
			case http.ErrServerClosed:
				return
			}
			slacklog.Trace(err)
		}
		return
	}

	logger.Info(appconst.AppName+" http server start", zap.String("address", server.Addr))
	if err = server.ListenAndServe(); err != nil {
		switch err {
		case http.ErrServerClosed:
			return
		}
		slacklog.Trace(err)
	}

	return
}

func run(
	ctx context.Context,
	address string,
	tls bool,
) (err error) {
	logger := applog.Logger()
	ctx = contextWithOSSignal(ctx, logger, os.Interrupt, os.Kill)
	var eg *errgroup.Group
	eg, ctx = errgroup.WithContext(ctx)
	server := newServer(flagAddress.String())

	defer func() {
		slacklog.LogAsync("stopped " + version.String())
	}()

	eg.Go(func() error {
		return listen(server, tls)
	})

	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return server.Shutdown(ctx)
		}
	})

	return eg.Wait()
}
