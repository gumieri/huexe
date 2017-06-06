package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
)

type LampState struct {
	On					bool		`json:"on"`
	Brightness	int			`json:"bri"`
	Alert				string	`json:"alert"`
	Reachable		bool		`json:"reachable"`
}

type Lamp struct {
	State	LampState	`json:"state"`
}

func baseUrl() (baseUrl string) {
	token := os.Getenv("HUE_TOKEN")
	address := os.Getenv("HUE_ADDRESS")

	return fmt.Sprintf("http://%s/api/%s", address, token)
}

func getLamp(lampId int) (lamp *Lamp, err error) {
	url := fmt.Sprintf("%s/lights/%d", baseUrl(), lampId)
	response, err := http.Get(url)

	if err != nil {
		return
	}

	defer response.Body.Close()

	lamp = new(Lamp)
	err = json.NewDecoder(response.Body).Decode(lamp)

	return
}

func putLamp(lampId int, lamp *Lamp) (err error) {
	url := fmt.Sprintf("%s/lights/%d/state", baseUrl(), lampId)

	marshal, err := json.Marshal(lamp)

	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(marshal))

	if err != nil {
		return
	}

	client := &http.Client{}

	_, err = client.Do(request)

	return
}

func main() {
	lampId := 1

	lamp, err := getLamp(lampId)

	if err != nil {
		log.Fatal(err)
		return
	}

	lamp.State.On = !lamp.State.On

	err = putLamp(lampId, lamp)

	if err != nil {
		log.Fatal(err)
	}
}
