# huexe

A group of executables to turn On / Off (`cmd/switch`) and Dim (`cmd/dimmer`) / Bright (`cmd/brighter`) a specific Philips Hue Lamp.

For now only the `cmd/switch` is implemented and the lampID is hardcoded as `1` just because of laziness.

# build

```bash
GOOS=windows GOARCH=amd64 go build --ldflags="-X main.address=192.168.1.2 -X main.token=username -H windowsgui" github.com/huexe/cmd/switch
GOOS=windows GOARCH=amd64 go build --ldflags="-X main.address=192.168.1.2 -X main.token=username -H windowsgui" github.com/huexe/cmd/dimmer
GOOS=windows GOARCH=amd64 go build --ldflags="-X main.address=192.168.1.2 -X main.token=username -H windowsgui" github.com/huexe/cmd/brighter
```
