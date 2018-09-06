package api

import (
	"io"
	"net/http"
	"strings"

	"github.com/tsaikd/KDGoLib/jsonex"
	"github.com/tsaikd/KDGoLib/structutil/reflectstruct"
	"github.com/tsaikd/KDGoLib/structutil/trimstructspace"
	"github.com/tsaikd/govalidator"
)

const (
	contentType = "Content-Type"
)

// BindParam bind param from http Request query, martini Params, http Request body, and validate
func BindParam(areq interface{}, req *http.Request) (err error) {
	reqQuery := req.URL.Query()
	for key, value := range reqQuery {
		if strings.HasSuffix(key, "[]") {
			reqQuery[strings.TrimSuffix(key, "[]")] = value
		}
	}

	if err = reflectstruct.ReflectStruct(areq, reqQuery); err != nil {
		return
	}

	if contentType := req.Header.Get(contentType); contentType != "" {
		if strings.Contains(contentType, "json") {
			if err = bindJSONBody(areq, req); err != nil {
				return
			}
		} else if strings.Contains(contentType, "multipart/form-data") {
			if err = bindMultipartForm(areq, req); err != nil {
				return
			}
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			if err = bindForm(areq, req); err != nil {
				return
			}
		}
	}

	if err = trimstructspace.TrimStructSpace(areq); err != nil {
		return
	}

	ok, err := govalidator.ValidateStruct(areq)
	if err != nil || !ok {
		return ErrInvalidParam.New(err)
	}

	return
}

func bindJSONBody(areq interface{}, req *http.Request) (err error) {
	if req.Body == nil {
		return
	}

	defer req.Body.Close()

	var value interface{}
	if err = jsonex.NewDecoder(req.Body).Decode(&value); err != nil && err != io.EOF {
		return
	}

	return reflectstruct.ReflectStruct(areq, value)
}

// 20 MB
const bindMaxMemory = 20 * 1024 * 1024

func bindMultipartForm(areq interface{}, req *http.Request) (err error) {
	if err = req.ParseMultipartForm(bindMaxMemory); err != nil {
		return
	}
	if err = reflectstruct.ReflectStruct(areq, req.MultipartForm.Value); err != nil {
		return
	}
	if err = reflectstruct.ReflectStruct(areq, req.MultipartForm.File); err != nil {
		return
	}
	return
}

func bindForm(areq interface{}, req *http.Request) (err error) {
	if err = req.ParseForm(); err != nil {
		return
	}
	if err = reflectstruct.ReflectStruct(areq, req.Form); err != nil {
		return
	}
	return
}
