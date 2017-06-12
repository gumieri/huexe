package main

import (
	"encoding/json"
	"github.com/andlabs/ui"
	"github.com/gumieri/huexe/lib/api"
	"io/ioutil"
	"net/http"
)

// NUPnP represents the response from the search of a Hue Bridge
type NUPnP struct {
	Address string `json:"internalipaddress"`
}

func setup() (err error) {
	config := new(api.Config)

	// Defaults
	config.Id = 1
	config.Address = ""
	config.Steps = 25

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

	keys := make([]NUPnP, 0)
	json.Unmarshal(responseBytes, &keys)

	if len(keys) > 0 {
		config.Address = keys[0].Address
	}

	err = api.SetConfig(config)

	if err != nil {
		return
	}

	return
}

func mainUi() {
	button := ui.NewButton("Setup")

	box := ui.NewVerticalBox()
	box.Append(button, false)

	window := ui.NewWindow("Setup", 200, 100, false)
	window.SetChild(box)

	button.OnClicked(func(*ui.Button) {
		setup()
		return
	})

	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	window.Show()
}

func main() {
	err := ui.Main(mainUi)

	if err != nil {
		panic(err)
	}
}
