package main

import (
	"flag"
	"log"

	auth "github.com/kyogai2281337/cns_eljur/internal/app/auth"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../../configs/auth.toml", "path to configs")
}

func main() {
	flag.Parse()

	config := auth.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := auth.Start(config); err != nil {
		log.Fatal(err)
	}
}
