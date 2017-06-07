# huexe

A group of executables to turn On / Off (`switch.go`) and Dim (`dimmer.go`) / Bright (`brighter.go`) a specific Philips Hue Lamp.

# build

```bash
GOOS=windows GOARCH=amd64 go build --ldflags="-X main.address=192.168.1.2 -X main.token=username -H windowsgui" switch.go
```
