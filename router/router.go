package router

import (
	"github.com/eBayClassifiedsGroup/ammonitrix/config"
	"github.com/gorilla/mux"
	"gopkg.in/olivere/elastic.v3"
)

//Elastic Connection
var elasticConf = &config.DefaultConfig.Elastic
var Elastic *elastic.Client

func NewRouter() *mux.Router {
	Elastic, _ = elastic.NewClient()
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
