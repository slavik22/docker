package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBSource            string        `mapstructure:"DB_SOURCE"`
	DBTestSource        string        `mapstructure:"DB_TEST_SOURCE"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	SecretKey           string        `mapstructure:"SECRET_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
