package router

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type RoutesV1 []Route

var routes = RoutesV1{
	Route{"Index", "GET", "/v1/", r.Index},
	Route{"DataIndex", "GET", "/v1/data", DataIndex},
	Route{"DataGet", "GET", "/v1/data/{dataId}", DataGet},
}
