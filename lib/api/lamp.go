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
// id refers to the Lamp ID from the API
// address refers to the bridge IP Address, the API Address
// username refers to a hash key/token generated from the API when registered a device/user
// It returns a Lamp with the current State
func GetLamp(id int, address string, username string) (lamp *Lamp, err error) {
	url := fmt.Sprintf("http://%s/api/%s/lights/%d", address, username, id)
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	lamp = new(Lamp)
	err = json.NewDecoder(response.Body).Decode(lamp)

	if err != nil {
		return nil, err
	}

	return
}

// PutLamp update the State of a Hue Lamp using the API
// id refers to the Lamp ID from the API
// address refers to the bridge IP Address, the API Address
// username refers to a hash key/token generated from the API when registered a device/user
func PutLamp(id int, address string, username string, lamp *Lamp) (err error) {
	url := fmt.Sprintf("http://%s/api/%s/lights/%d/state", address, username, id)

	marshal, err := json.Marshal(lamp.State)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(marshal))

	if err != nil {
		return err
	}

	client := &http.Client{}

	_, err = client.Do(request)

	if err != nil {
		return err
	}

	return
}
