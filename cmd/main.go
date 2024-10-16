package main

import (
	"github.com/hritesh04/synlabs/api"
	"github.com/hritesh04/synlabs/config"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		panic(err)
	}
	api.SetupServer(cfg)
}
