package server

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/tsaikd/full-stack-demo/server/appconst"
	"github.com/tsaikd/full-stack-demo/server/applog"
	"github.com/tsaikd/full-stack-demo/server/slacklog"
	"go.uber.org/zap"
)

var webRoot = filepath.Join(appconst.ProjectRootDir, "web", "dist")
var webFileServer = http.FileServer(http.Dir(webRoot))

var webIndexContent = []byte{}
var webIndexMetaContent = []byte{}
var webIndexModified = time.Time{}

func initWebIndex() ([]byte, []byte, time.Time) {
	webIndex := filepath.Join(webRoot, "index.html")
	fi, err := os.Stat(webIndex)
	if err != nil {
		if os.IsNotExist(err) {
			applog.Logger().Warn("no web index found")
		} else {
			slacklog.Trace(err)
		}
		return []byte{}, []byte{}, time.Time{}
	}
	data, err := ioutil.ReadFile(webIndex)
	if err != nil {
		slacklog.Trace(err)
		return []byte{}, []byte{}, time.Time{}
	}
	return data, []byte{}, fi.ModTime()
}

var webJSContent = io.ReadSeeker(bytes.NewReader([]byte{}))
var webJSModified = time.Time{}

func initWebJS() (io.ReadSeeker, time.Time) {
	files, err := filepath.Glob(filepath.Join(webRoot, "js", "index.*.js"))
	if err != nil {
		slacklog.Trace(err)
		return nil, time.Time{}
	}
	if len(files) < 1 {
		applog.Logger().Warn("no web javascript found")
		return nil, time.Time{}
	}
	if len(files) > 1 {
		applog.Logger().Warn("find more than one web javascript, use first file", zap.Strings("files", files))
	}

	body := &bytes.Buffer{}
	fpath := files[0]
	fi, err := os.Stat(fpath)
	if err != nil {
		slacklog.Trace(err)
		return nil, time.Time{}
	}
	modified := fi.ModTime()

	fp, err := os.Open(fpath)
	if err != nil {
		slacklog.Trace(err)
		return nil, time.Time{}
	}
	if _, err = io.Copy(body, fp); err != nil {
		slacklog.Trace(err)
		return nil, time.Time{}
	}
	if err = fp.Close(); err != nil {
		slacklog.Trace(err)
		return nil, time.Time{}
	}

	return bytes.NewReader(body.Bytes()), modified
}
