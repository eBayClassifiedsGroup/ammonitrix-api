package router

import (
	"log"
	"net/http"
)

//Route handles API Routing
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//RoutesV1 is an Array of routes
type RoutesV1 []Route

var routes = RoutesV1{
	Route{"Root", "GET", "/", Root},
	Route{"Index", "GET", "/v1/", GetIndex},
	Route{"DataIndex", "GET", "/v1/data", GetDataIndex},
	Route{"DataGet", "GET", "/v1/data/{dataID}", GetData},
}

func Tracer(method, url, body string) {
	log.Printf("Requesting %s %s", method, url)
	log.Printf("Request body: %s", body)
}
