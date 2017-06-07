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

type LampState struct {
	On bool `json:"on"`
}

type Lamp struct {
	State LampState `json:"state"`
}

func getLamp(lampId int) (lamp *Lamp, err error) {
	url := fmt.Sprintf("http://%s/api/%s/lights/%d", address, token, lampId)
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
	url := fmt.Sprintf("http://%s/api/%s/lights/%d/state", address, token, lampId)

	marshal, err := json.Marshal(lamp.State)

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
