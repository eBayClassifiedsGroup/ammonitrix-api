package router

import (
	"github.com/eBayClassifiedsGroup/ammonitrix/config"
	"github.com/gorilla/mux"
	"github.com/mattbaird/elastigo/lib"
)

//Elastic Connection
var elasticConf = &config.DefaultConfig.Elastic
var Elastic = elastigo.NewConn()

func NewRouter() *mux.Router {
	Elastic.Domain = elasticConf.Host
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
