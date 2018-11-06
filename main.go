package main

import (
	"log"
	"net/http"

	"github.com/samuka182/archer/routes"
	util "github.com/samuka182/archer/utils"
)

func main() {

	routerArchttp := util.NewRouter(routes.RoutesArchttp)

	log.Fatal(http.ListenAndServe(":8080", routerArchttp))
}
