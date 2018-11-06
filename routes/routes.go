package routes

import (
	"github.com/samuka182/archer/handlers"
	"github.com/samuka182/archer/structs"
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
	structs.Route{
		"CreateArchestration",
		"POST",
		"/archestration",
		handlers.CreateArchestration,
	},
	structs.Route{
		"GetArchestration",
		"GET",
		"/archestration/{id}",
		handlers.GetArchestration,
	},
	structs.Route{
		"GetArchestrations",
		"GET",
		"/archestrations",
		handlers.GetArchestrations,
	},
}
