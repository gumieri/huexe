package api

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Config contain the necessary data to interact with the API
type Config struct {
	Id       int
	Address  string
	Username string
}

// GetConfig read the config.yml file in the executable directory and return it
func GetConfig() (config Config, err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(dir)

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	config.Address = viper.GetString("address")
	config.Username = viper.GetString("username")
	config.Id = viper.GetInt("id")

	return
}
