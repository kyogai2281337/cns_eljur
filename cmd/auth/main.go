package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kyogai2281337/cns_eljur/internal/auth/controller"
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
	log.Println("[INFO] => starting server using ", config.DatabaseURL, " URL: ")
	if err := controller.Start(config); err != nil {
		log.Fatal(err)
	}
	//time.Sleep(10*time.Second)
}
