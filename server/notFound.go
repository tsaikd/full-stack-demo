package server

import (
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/tsaikd/KDGoLib/webutil/render/httpRender"
	"github.com/tsaikd/full-stack-demo/server/api"
)

const cacheHeaderOneYear = "max-age=31536000"

type notFoundHandler struct {
}

func (t notFoundHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	fileServer := webFileServer

	upath := req.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		req.URL.Path = upath
	}

	name := path.Clean(upath)
	if isAPIPath(name) {
		api.ProvideError(w, req, api.LogTypeNameAPI, start, api.ErrAPINotFound.New(nil))
		return
	}

	if isCacheablePath(name) {
		w.Header().Set(httpRender.HeaderCacheControl, cacheHeaderOneYear)
	}

	api.ProvideHandler(w, req, api.LogTypeNamePage, fileServer.ServeHTTP, start)
	return
}

var reAPIPath = regexp.MustCompile(`(?i)^/api/`)

func isAPIPath(fpath string) bool {
	return reAPIPath.MatchString(fpath)
}

var reResourceFile = regexp.MustCompile(`(?i)\.(js|css|map|png|ico|gif|jpg|jpeg|svg)$`)

func isResourcePath(fpath string) bool {
	return reResourceFile.MatchString(fpath)
}

var reCacheablePath = regexp.MustCompile(`(?i)^/static/.+\.[0-9a-f]{6,}\.(js|css|map|png|ico|gif|jpg|jpeg|svg)$`)

func isCacheablePath(fpath string) bool {
	return reCacheablePath.MatchString(fpath)
}

func maxTime(times ...time.Time) time.Time {
	result := time.Time{}
	for _, t := range times {
		if t.After(result) {
			result = t
		}
	}
	return result
}
