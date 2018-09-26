package main

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("AGENDA_API")
}

type config struct {
	Port          int    `mapstructure:"port"`
	MongoHost     string `mapstructure:"mongo_host"`
	MongoPort     int    `mapstructure:"mongo_port"`
	MongoUsername string `mapstructure:"mongo_username"`
	MongoPassword string `mapstructure:"mongo_password"`
}

func unmarshalConfig() (*config, error) {
	var c config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
