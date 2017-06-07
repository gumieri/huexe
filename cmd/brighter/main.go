package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type LampState struct {
	Brightness int `json:"bri"`
}

type Lamp struct {
	State LampState `json:"state"`
}

func getLamp(id int, address string, username string) (lamp *Lamp, err error) {
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

func putLamp(id int, address string, username string, lamp *Lamp) (err error) {
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

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	address := viper.GetString("address")
	username := viper.GetString("username")
	id := viper.GetInt("id")

	lamp, err := getLamp(id, address, username)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(lamp)
	lamp.State.Brightness = lamp.State.Brightness + 25

	if lamp.State.Brightness > 254 {
		lamp.State.Brightness = 254
	}

	err = putLamp(id, address, username, lamp)

	if err != nil {
		log.Fatal(err)
		return
	}
}
