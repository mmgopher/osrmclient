package service

import (
	"shipping/logging"
	"shipping/model"
	"sort"
	"sync"
)

type RoutesCalculator struct {
	pathsRouter PathRouter
}

func NewRoutesCalculator(pr PathRouter) *RoutesCalculator {
	return &RoutesCalculator{pr}
}

func (rc RoutesCalculator) CalculateRoutes(source string, destinations []string) ([]model.Route) {

	routesChannel := make(chan model.Route, len(destinations))
	var wg sync.WaitGroup
	for _, destination := range destinations {
		wg.Add(1)
		go getRouteDetails(source, destination, routesChannel, &wg, rc.pathsRouter)
	}
	wg.Wait()
	close(routesChannel)
	routeList := []model.Route{}
	for elem := range routesChannel {
		routeList = append(routeList, elem)
	}
	sort.Sort(model.ByTime(routeList))
	return routeList
}

func getRouteDetails(source string, destination string, routesChannel chan <-model.Route, wg *sync.WaitGroup, pathsRouter PathRouter) {
	defer wg.Done()
	logging.Info("Calculate route for destination", destination)
	route, error := pathsRouter.GetRoute(source, destination)
	if (error != nil) {
		logging.Error(error)
	} else {
		routesChannel <- route
	}
}