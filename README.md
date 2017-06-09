# Huexe

A group of executables to turn On / Off (`switch`) and Dim (`dimmer`) / Bright (`brighter`) a specific Philips Hue Lamp.

## Configurations

The executables read a `config.yml` in the working directory.

Example of config:

```yml
id: 1
address: 192.168.1.2
username: 249ba36000029bbe97499c03db5a9001f6b734ec
step_size: 25
```
### How to get the `username`

- https://www.developers.meethue.com/documentation/getting-started

## Code Status

[![Go Report Card](https://goreportcard.com/badge/github.com/gumieri/huexe)](https://goreportcard.com/report/github.com/gumieri/huexe)
[![Build Status](https://travis-ci.org/gumieri/huexe.svg?branch=master)](https://travis-ci.org/gumieri/huexe)


## License

Huexe is released under the [MIT License](http://www.opensource.org/licenses/MIT).
