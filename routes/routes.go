package routes

import (
	handlers "archer/handlers"
	structs "archer/structs"
)

//Routes ...
var RoutesArchttp = []structs.Route{
	structs.Route{
		"Request",
		"POST",
		"/request",
		handlers.FullRequest,
	},
}
