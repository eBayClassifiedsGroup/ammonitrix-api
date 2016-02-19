package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/build"
	"github.com/eBayClassifiedsGroup/ammonitrix-api/config"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/router"
)

//Version number
var Version = build.Version

func main() {
	var filename string
	var v bool

	flag.StringVar(&filename, "cfg", "", "path to config file")
	flag.BoolVar(&v, "v", false, "show version")
	flag.Parse()

	if v {
		fmt.Println(Version)
		return
	}
	log.Printf("[INFO] Version %s starting", Version)

	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(config.DefaultConfig.Listen.Port, router))
}
