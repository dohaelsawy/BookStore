package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSource     string `mapstructure:"DB_SOURCE"`
	DBDriver     string `mapstructure:"DB_DRIVER"`
	ServerAdress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadAppEnv(path string) (Config, error) {
	config := Config{}

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv() //AutomaticEnv makes Viper check if environment variables match any of the existing keys (config, default or flags). If matching env vars are found, they are loaded into Viper.
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
