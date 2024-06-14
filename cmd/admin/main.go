package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kyogai2281337/cns_eljur/internal/admin/controller"
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

	if err := controller.Start(config); err != nil {
		log.Fatal(err)
	}
}
