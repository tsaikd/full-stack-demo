package api

import (
	"net/http"
)

// API contains info for route resource
type API struct {
	Method  Method
	Pattern string
	Handler func(w http.ResponseWriter, req *http.Request)
}

var registedAPIs = []*API{}

// All return all registed APIs
func All() []*API {
	return registedAPIs
}

// Add API module
func Add(module *API) {
	registedAPIs = append(registedAPIs, module)
}
