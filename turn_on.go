package main

import (
  "bytes"
  "fmt"
  "log"
  "os"
  "net/http"
)

func main() {
  token := os.Getenv("HUE_TOKEN")
  address := os.Getenv("HUE_ADDRESS")
  lamp := os.Getenv("HUE_LAMP")

  url := fmt.Sprintf("http://%s/api/%s/lights/%s/state", address, token, lamp)

  dataJson := []byte(`{ "on":"true" }`)

  request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(dataJson))
  if err != nil {
    log.Fatal(err)
  }

  client := &http.Client{}

  _, err = client.Do(request)
  if err != nil {
    log.Fatal(err)
  }
}
