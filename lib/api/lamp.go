package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// LampState represents the State of a Hue Lamp from the API
type LampState struct {
	On         bool `json:"on"`
	Brightness int  `json:"bri"`
	Reachable  bool `json:"reachable"`
}

// Lamp hold informations of a Hue Lamp from the API
type Lamp struct {
	State LampState `json:"state"`
}

// GetLamp request to the API for informations about a Hue Lamp
// It returns a Lamp with the current State
func GetLamp(config Config) (lamp *Lamp, err error) {
	id := config.Id
	address := config.Address
	username := config.Username

	url := fmt.Sprintf("http://%s/api/%s/lights/%d", address, username, id)
	response, err := http.Get(url)

	if err != nil {
		return
	}

	defer response.Body.Close()

	lamp = new(Lamp)
	err = json.NewDecoder(response.Body).Decode(lamp)

	if err != nil {
		return
	}

	return
}

// PutLamp update the State of a Hue Lamp using the API
func PutLamp(config Config, lamp *Lamp) (err error) {
	marshal, err := json.Marshal(lamp.State)

	if err != nil {
		return
	}

	id := config.Id
	address := config.Address
	username := config.Username

	url := fmt.Sprintf("http://%s/api/%s/lights/%d/state", address, username, id)

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(marshal))

	if err != nil {
		return
	}

	client := &http.Client{}

	_, err = client.Do(request)

	if err != nil {
		return
	}

	return
}
