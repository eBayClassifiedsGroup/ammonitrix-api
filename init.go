package main

import (
	"log"

	"github.com/eBayClassifiedsGroup/ammonitrix-api/config"
)

func loadConfig(filename string) *config.Config {
	cfg, err := config.FromFile(filename)
	if err != nil {
		log.Fatal("[FATAL] ", err)
	}
	return cfg
}
