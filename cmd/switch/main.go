package main

import (
	"github.com/gumieri/huexe/lib/api"
	"log"
)

func main() {
	var err error

	config, err := api.GetConfig()

	lamp, err := api.GetLamp(config)

	if err != nil {
		log.Fatal(err)
		return
	}

	lamp.State.On = !lamp.State.On

	err = api.PutLamp(config, lamp)

	if err != nil {
		log.Fatal(err)
		return
	}
}
