package main

import (
	"encoding/json"
	"github.com/gumieri/huexe/lib/api"
	"io/ioutil"
	"log"
	"net/http"
)

type NUPnP struct {
	Address string `json:"internalipaddress"`
}

func main() {
	var err error

	url := "https://www.meethue.com/api/nupnp"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	keys := make([]NUPnP, 0)
	json.Unmarshal(responseBytes, &keys)

	config := new(api.Config)

	// Defaults
	config.Id = 1
	config.Address = keys[0].Address
	config.Steps = 25

	err = api.SetConfig(config)

	if err != nil {
		log.Fatal(err)
	}
}
