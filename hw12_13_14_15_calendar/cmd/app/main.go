package main

import (
	"flag"
	"log"

	"github.com/mxmntv/otus_hw/hw12_calendar/config"
	"github.com/mxmntv/otus_hw/hw12_calendar/internal/app"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./../config/config.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
