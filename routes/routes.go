package routes

import (
	handlers "archer/handlers"
	structs "archer/structs"
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
