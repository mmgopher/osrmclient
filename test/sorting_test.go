package test

import (
	"testing"
	"shipping/model"
	"sort"
	"fmt"
	"reflect"
)

func TestRoutesSorting(t *testing.T) {

	r1 := model.Route{Destination:"13.397634,52.529407", Duration:412.74, Distance:456.7}
	r2 := model.Route{Destination:"13.397634,52.529407", Duration: 411.7, Distance: 353.6}
	r3 := model.Route{Destination:"13.397634,52.529407", Duration: 411.7, Distance: 356.6}
	r4 := model.Route{Destination:"13.397634,52.529407", Duration: 126, Distance: 111}
	expected := []model.Route{r4, r2, r3, r1}
	routeList := []model.Route{r1, r2, r3, r4}
	testSorting(t, expected, routeList)

	routeList = []model.Route{r4, r3, r2, r1}
	testSorting(t, expected, routeList)

	routeList = []model.Route{r4, r2, r3, r1}
	testSorting(t, expected, routeList)

}

func testSorting(t *testing.T, expected []model.Route, routeList []model.Route) {
	sort.Sort(model.ByTime(routeList))

	if (!reflect.DeepEqual(expected, routeList)) {
		t.Errorf("Wrong sorting result \n expected: %s \n actual: %s", fmt.Sprint(expected), fmt.Sprint(routeList))
	}

}