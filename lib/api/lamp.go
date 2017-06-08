package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LampState struct {
	On         bool `json:"on"`
	Brightness int  `json:"bri"`
	Reachable  bool `json:"reachable"`
}

type Lamp struct {
	State LampState `json:"state"`
}

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
