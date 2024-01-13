package config

import (
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./../")

	err := config.ReadInConfig()

	if err != nil {
		//panic(fmt.Errorf("Fatal error config file: %w \n", err))
		panic(err)
	}

	return config
}
