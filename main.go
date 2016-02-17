package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/config"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/router"
)

var version = "0.0.1-dev"

func main() {
	var filename string
	var v bool

	flag.StringVar(&filename, "cfg", "", "path to config file")
	flag.BoolVar(&v, "v", false, "show version")
	flag.Parse()

	if v {
		fmt.Println(version)
		return
	}
	log.Printf("[INFO] Version %s starting", version)

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(config.DefaultConfig.Listen.Port, router))
}
