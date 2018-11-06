package util

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samuka182/archer/structs"
)

//NewRouter ...
func NewRouter(routes []structs.Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
