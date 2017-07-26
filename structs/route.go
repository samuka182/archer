package structs

import "net/http"

//Route represents a Route data
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
