package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Password string `mapstructure:PASSWORD`
}

func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.ReadInConfig()

	viper.Unmarshal(&config)

	return 
} 