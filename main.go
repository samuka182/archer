package main

import (
	"archer/routes"
	util "archer/utils"
	"log"
	"net/http"
)

func main() {

	routerArchttp := util.NewRouter(routes.RoutesArchttp)

	log.Fatal(http.ListenAndServe(":8080", routerArchttp))
}
