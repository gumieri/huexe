package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// NUPnP represents the response from the search of a Hue Bridge
type NUPnP struct {
	Address string `json:"internalipaddress"`
}

func DiscoverBridgeAddress() (address string, err error) {
	url := "https://www.meethue.com/api/nupnp"
	response, err := http.Get(url)

	if err != nil {
		return
	}

	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	bridges := make([]NUPnP, 0)
	json.Unmarshal(responseBytes, &bridges)

	if len(bridges) > 0 {
		address = bridges[0].Address
	}

	return
}
