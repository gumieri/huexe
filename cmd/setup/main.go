package main

import (
	"github.com/gumieri/huexe/lib/api"
	"log"
)

func main() {
	config := new(api.Config)

	// Defaults
	config.Id = 1
	config.Address = ""
	config.Steps = 25

	bridgeAddress, err := api.DiscoverBridgeAddress()

	if err != nil {
		return
	}

	config.Address = bridgeAddress

	username, err := api.RegisterUser(*config)

	if err != nil {
		log.Fatal(err)
		return
	}

	config.Username = username

	err = api.SetConfig(config)

	if err != nil {
		log.Fatal(err)
	}
}
