package osrm

import (
	"shipping/model"
	"net/http"
	"io/ioutil"
	"bytes"
	"shipping/logging"
	"encoding/json"
	"errors"
	"fmt"
)

type PathRoute struct {

}

func NewPathRoute() *PathRoute {
	return &PathRoute{}
}

func (PathRoute) GetRoute(src string, dst string) (model.Route, error) {
	var buffer bytes.Buffer
	buffer.WriteString("http://router.project-osrm.org/route/v1/driving/")
	buffer.WriteString(src)
	buffer.WriteString(";")
	buffer.WriteString(dst)
	req, err := http.NewRequest("GET", buffer.String(), nil)
	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()
	q.Add("overview", "false")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		logging.Error(err)
		return model.Route{}, err
	}

	req.Close = true

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		logging.Error(err)
		return model.Route{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Error(err)
		return model.Route{}, err
	}

	var resultMap map[string]interface{}
	json.Unmarshal(body, &resultMap)
	code := resultMap["code"];
	if code == "Ok" {
		routes := resultMap["routes"].([]interface{});
		route := routes[0].(map[string]interface{});
		duration := route["duration"].(float64)
		distance := route["distance"].(float64)
		return model.Route{Destination: dst, Duration: duration, Distance: distance}, nil
	} else {
		return model.Route{}, errors.New(fmt.Sprintf("http: %s", code))
	}

}