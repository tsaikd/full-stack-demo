package webclient

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/tsaikd/KDGoLib/version"
	"github.com/tsaikd/KDGoLib/webutil"
	"github.com/tsaikd/full-stack-demo/server/appconst"
)

var client = &http.Client{}
var userAgent = appconst.AppName + "_Bot/" + version.VERSION

// NewRequest returns a new Request given a method, URL, and optional body.
func NewRequest(method string, url string, body io.Reader) (req *http.Request, err error) {
	if req, err = http.NewRequest(method, url, body); err != nil {
		return
	}

	req.Header.Set("User-Agent", userAgent)
	return
}

// Get issues a GET to the specified URL
func Get(url string) (resp *http.Response, err error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	if resp, err = client.Do(req); err != nil {
		return
	}
	err = webutil.NewResponseError(resp)
	return
}

// GetJSON issues a GET to the specified URL and unmarshal response
func GetJSON(url string, data interface{}) (resp *http.Response, err error) {
	if resp, err = Get(url); err != nil {
		return
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(data)
	return
}

// GetString issues a GET to the specified URL and read body as string
func GetString(url string) (body string, resp *http.Response, err error) {
	if resp, err = Get(url); err != nil {
		return
	}

	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(raw), resp, err
}
