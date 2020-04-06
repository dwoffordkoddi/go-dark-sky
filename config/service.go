package config

import (
	"fmt"
	"github.com/KoddiDev/koddi-framework/server/environment"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	APP     app     `yaml:"app"`
	DarkSky DarkSky `yaml:"darkSky"`
}

type app struct {
	Version string `yaml:"version"`
}

type DarkSky struct {
	Url    string `yaml:"url"`
	ApiKey string `yaml:"apiKey"`
}

var (
	ServiceConf *ServiceConfig //exported Configuration that all packages can use
)

/*
	LoadConfig will try to load a configuration file for the given repo
	The conf file should be named "service.yaml" in the "config" dir
*/
func LoadConfig() (*ServiceConfig, error) {
	//Get the environment to run in...
	env, _ := environment.GetEnvironment()
	configFile := fmt.Sprintf("service-%s", env)

	viperInstance := viper.New()
	viperInstance.AutomaticEnv()
	viperInstance.AddConfigPath("config")
	viperInstance.SetConfigName(configFile)

	//try to load the config file for the environment
	if err := viperInstance.ReadInConfig(); err != nil {
		return nil, errors.Wrapf(err, "Failed to read configuration file")
	}

	ServiceConf = &ServiceConfig{}
	if err := viperInstance.Unmarshal(ServiceConf); err != nil {
		return nil, errors.Wrapf(err, "Failed to parse configuration")
	}

	return ServiceConf, nil
}
