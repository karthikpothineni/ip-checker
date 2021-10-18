package config

import (
	"github.com/spf13/viper"

	"ip-checker/logger"
)

var config *viper.Viper

// Init - exported method that starts the viper and returns the configuration struct.
func Init(path string) {
	var err error
	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath(path)
	v.SetConfigName("config")
	err = v.MergeInConfig()
	if err != nil {
		logger.Log.Fatal("Error on parsing configuration file. Error " + err.Error())
	}
	config = v
}

// GetConfig - function to expose the config object
func GetConfig() *viper.Viper {
	return config
}
