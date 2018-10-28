package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"shipping/logging"
	"shipping/model"
	"encoding/json"
	"shipping/service"
)

type RoutesController struct {
	routeCalculator *service.RoutesCalculator
}


func NewRoutesController(rc *service.RoutesCalculator) *RoutesController {
	return &RoutesController{rc}
}


func (rc RoutesController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryValues := r.URL.Query()
	src := queryValues.Get("src")
	if(src == "" || len(queryValues["dst"]) == 0) {
		responseWithError(w, http.StatusBadRequest, "Invalid request data")
	} else {
		routeList := rc.routeCalculator.CalculateRoutes(src, queryValues["dst"])
		response := model.Response{src, routeList}
		responseWithJson(w, http.StatusOK, response)
	}
}

func responseWithError(w http.ResponseWriter, code int,  msg string) {
	responseWithJson(w, code, map[string]string{"error": msg})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if(err != nil) {
		logging.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}