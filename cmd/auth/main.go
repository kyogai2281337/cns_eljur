package main

import (
	"flag"
	"log"

	auth "github.com/kyogai2281337/cns_eljur/internal/auth/auth"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./auth.toml", "path to configs")
}

func main() {
	flag.Parse()

	config := auth.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[INFO] => starting server using ", config.DatabaseURL, " URL: ")
	if err := auth.Start(config); err != nil {
		log.Fatal(err)
	}
	//time.Sleep(10*time.Second)
}
