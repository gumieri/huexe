package main

import (
	"github.com/gumieri/huexe/lib/api"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	address := viper.GetString("address")
	username := viper.GetString("username")
	id := viper.GetInt("id")

	lamp, err := api.GetLamp(id, address, username)

	if err != nil {
		log.Fatal(err)
		return
	}

	lamp.State.On = !lamp.State.On

	err = api.PutLamp(id, address, username, lamp)

	if err != nil {
		log.Fatal(err)
		return
	}
}
