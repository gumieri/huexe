# Huexe

A group of executables to turn On / Off (`switch`) and Dim (`dimmer`) / Bright (`brighter`) a specific Philips Hue Lamp.

## Installation

Download the package of your Operating System at [hue/releases](https://github.com/gumieri/huexe/releases) and extract it where you prefer.

Use the `setup` executable to generate a `config.yml`.

It will search for the Philips Hue Bridge in your network using [meethue.com/api/nupnp](https://www.meethue.com/api/nupnp).

All the others executables use this same `config.yml`.

You will need to open this file in a text editor and configure manualy the rest.

If you want to control other lamps, create a copy of this instalation and edit the `id` to the desired one.
### Configurations

#### id
Refers to the Lamp ID from the Philips Hue Bridge API.
When requesting the API for all Lamps it's the `key` refering the JSON object, with `state`, `type`, etc..

It's is **not** the `uniqueid`, `modelid` nor `luminaireuniqueid`.
#### address
The Philips Hue Bridge IP Address.

There is many ways to get it, one way is accessing [meethue.com/api/nupnp](https://www.meethue.com/api/nupnp). It wil show the `address` as `internalipaddress`:
```json
[{"id":"01890375a60c8f","internalipaddress":"192.168.1.2"}]
```
#### username
It's a Hash Key / Token generated from the Philips Hue Bridge API when registered a Device / User.

Follow the [official Getting Started Guide](https://www.developers.meethue.com/documentation/getting-started) to get the username.
#### steps
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
