package main

import (
	"github.com/gumieri/huexe/lib/api"
	"log"
)

func main() {
	var err error

	config, err := api.GetConfig()

	if err != nil {
		log.Fatal(err)
		return
	}

	err = api.UnregisterUser(config)

	if err != nil {
		log.Fatal(err)
	}
}
