# Huexe

A group of executables to turn On / Off (`switch`) and Dim (`dimmer`) / Bright (`brighter`) a specific Philips Hue Lamp.

## Installation

1. Download the package of your Operating System at [huexe/releases](https://github.com/gumieri/huexe/releases) and extract it where you prefer.

2. Press the Philips Hue Bridge "Push-Link Button" and (by the limit of 30 seconds) execute the `setup`.
    The Huexe will be registered and the Key will be stored at `config.yml` file.

It will be configured to use the lamp with `id` `1`. If you want to control other lamps, create a copy of this installation and edit the `id` inside `config.yml` to the desired one.

## Configurations

### id
Refers to the Lamp ID from the Philips Hue Bridge API.

If you are using the API to find the lamps IDs, it is **not** the `uniqueid`, `modelid` nor `luminaireuniqueid`.
It's the `key` refering the JSON object, with `state`, `type`, etc..
### address
The Philips Hue Bridge IP Address.
### username
It's a Hash Key / Token generated from the Philips Hue Bridge API when the Huexe is registered.
### steps
Refers to how much one execution of `dimmer` / `brighter` will dim / bright.

The Brightness of a Hue lamp goes from 1 to 254, and it just accept integer values. In this way the default value of `steps` is `25`, to a total of 10 executions from minimum to maximum bright.

### Example
```yml
id: 1
address: 192.168.1.2
username: 249ba36000029bbe97499c03db5a9001f6b734ec
steps: 25
```
## Code Status

[![Go Report Card](https://goreportcard.com/badge/github.com/gumieri/huexe)](https://goreportcard.com/report/github.com/gumieri/huexe)
[![Build Status](https://travis-ci.org/gumieri/huexe.svg?branch=master)](https://travis-ci.org/gumieri/huexe)


## License

Huexe is released under the [MIT License](http://www.opensource.org/licenses/MIT).
