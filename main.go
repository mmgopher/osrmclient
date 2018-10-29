package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"shipping/config"
	"shipping/service"
	"shipping/service/osrm"
	"shipping/controllers"
)

func main() {
	pathRouter := osrm.NewPathRoute()
	routeCalculator := service.NewRoutesCalculator(pathRouter)
	routesController :=controllers.NewRoutesController(routeCalculator)
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/routes", routesController.Get)
	log.Fatal(http.ListenAndServe("localhost:80", router))
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config.Template.ExecuteTemplate(w, "index.gohtml", nil)
}
