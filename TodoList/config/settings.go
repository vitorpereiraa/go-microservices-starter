package config

import (
	"github.com/spf13/viper"
)

type Settings struct {
	SERVER_PORT string `mapstructure:"SERVER_PORT"`
	SERVER_HOST string `mapstructure:"SERVER_HOST"`

	DB_HOST string `mapstructure:"DB_HOST"`
	DB_DRIVER string `mapstructure:"DB_DRIVER"`
	DB_USER string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_NAME string `mapstructure:"DB_NAME"`
	DB_PORT string `mapstructure:"DB_PORT"`
}

func LoadSettings() (s Settings, err error) {
	vp := viper.New()

	vp.SetConfigFile(".env")
	vp.AddConfigPath(".")

	vp.AutomaticEnv()

	if err := vp.ReadInConfig(); err != nil {
		return s, err
	}

	if err := vp.Unmarshal(&s); err != nil {
		return s, err
	}

	return s, nil 
}
