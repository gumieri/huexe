package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User is used to register a new username
type User struct {
	DeviceType string `json:"devicetype"`
}

// RegisterError represent the data from a failed RegisterResponse
type RegisterError struct {
	ErrorType   int    `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

// RegisterSuccess represent the data from a succeeded RegisterResponse
type RegisterSuccess struct {
	Username string `json:"username"`
}

// RegisterResponse represent the response from registering a username
type RegisterResponse struct {
	ErrorData   RegisterError   `json:"error"`
	SuccessData RegisterSuccess `json:"success"`
}

// RegisterUser ask for Philips Hue Bridge to register a new username
func RegisterUser(config Config) (username string, err error) {
	user := new(User)
	user.DeviceType = "Huexe"

	marshal, err := json.Marshal(user)

	if err != nil {
		return
	}

	url := fmt.Sprintf("http://%s/api", config.Address)
	payload := bytes.NewBuffer(marshal)
	request, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return
	}

	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	registers := make([]RegisterResponse, 0)
	json.Unmarshal(responseBytes, &registers)

	username = registers[0].SuccessData.Username

	return
}

// UnregisterUser delete the username on config from the Philips Hue Bridge
func UnregisterUser(config Config) (err error) {
	urlS := "http://%s/api/%s/config/whitelist/%s"
	url := fmt.Sprintf(urlS, config.Address, config.Username, config.Username)
	request, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return
	}

	client := &http.Client{}

	_, err = client.Do(request)

	return
}
