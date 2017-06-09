package api

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config contain the necessary data to interact with the API
type Config struct {
	Id       int    `yaml:id`
	Address  string `yaml:address`
	Username string `yaml:username`
	Steps    int    `yaml:steps`
}

func getConfigPath() (configPath string, err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return
	}

	configPath = filepath.Join(dir, "config.yml")

	return
}

// GetConfig read the config.yml file in the executable directory and return it
func GetConfig() (config Config, err error) {
	configPath, err := getConfigPath()

	if err != nil {
		return
	}

	configBytes, err := ioutil.ReadFile(configPath)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(configBytes, &config)

	if err != nil {
		return
	}

	return
}

func SetConfig(config *Config) (err error) {
	configPath, err := getConfigPath()

	if err != nil {
		return
	}

	configYaml, err := yaml.Marshal(config)

	if err != nil {
		return
	}

	ioutil.WriteFile(configPath, configYaml, 0640)

	if err != nil {
		return
	}

	return
}
