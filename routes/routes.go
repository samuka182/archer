package routes

import (
	"github.com/samuka182/archer/handlers"
	structs "github.com/samuka182/archer/structs"
)

//RoutesArchttp ...
var RoutesArchttp = []structs.Route{
	structs.Route{
		"Request",
		"POST",
		"/request",
		handlers.FullRequest,
	},
	structs.Route{
		"Get",
		"GET",
		"/get",
		handlers.GetTest,
	},
}
