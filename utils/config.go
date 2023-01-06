package utils

import (
	"github.com/moonman/mbank/global"
	"github.com/spf13/viper"
)

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		return err
	}

	return nil
}
