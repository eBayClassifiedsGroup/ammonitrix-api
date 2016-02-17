package router

import (
	"github.com/eBayClassifiedsGroup/ammonitrix/backends/elastic"
	"github.com/eBayClassifiedsGroup/ammonitrix/config"
	"github.com/gorilla/mux"
)

type Receiver struct {
	Config   *config.Config
	Elastic  *elastic.Elastic
	Metadata map[string]config.ElasticMetadata
}

func NewReceiver(config *config.Config) (*Receiver, error) {
	r := &Receiver{
		Config: config,
	}
	return r, nil
}

func (r *Receiver) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
