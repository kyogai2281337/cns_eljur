package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kyogai2281337/cns_eljur/internal/adminPanel/adminPanel_controller"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./config.toml", "path to configs")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := adminPanel_controller.Start(config); err != nil {
		log.Fatal(err)
	}
}
