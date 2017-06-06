package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var address string
var token string

type LampPutState struct {
	On bool `json:"on"`
}

type LampState struct {
	On         bool   `json:"on"`
	Brightness int    `json:"bri"`
	Alert      string `json:"alert"`
	Reachable  bool   `json:"reachable"`
}

type Lamp struct {
	State LampState `json:"state"`
}

func baseUrl() (baseUrl string) {
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

	lampPutState := new(LampPutState)
	lampPutState.On = lamp.State.On

	marshal, err := json.Marshal(lampPutState)

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
