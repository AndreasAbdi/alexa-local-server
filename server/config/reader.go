package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const configName = ".serverconf"

const keyServerAddress = "serverAddress"
const keyAlexaAppID = "alexaAppID"
const keyGoogleKey = "googleKey"

var paths = []string{
	"./.als",
	"$HOME/.als",
	".",
}

//Wrapper object of read in config.
type Wrapper struct {
	GoogleKey     string //key for accessing youtube/other google apis
	AlexaAppID    string //id for the alexa application.
	ServerAddress string //address for the server to deploy to.
}

//GetConfig object from system/default values.
func GetConfig() Wrapper {
	config := viper.New()
	config.SetConfigName(configName)
	for _, path := range paths {
		config.AddConfigPath(path)
	}

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return Wrapper{
		GoogleKey:     config.GetString(keyGoogleKey),
		AlexaAppID:    config.GetString(keyAlexaAppID),
		ServerAddress: config.GetString(keyServerAddress),
	}
}
