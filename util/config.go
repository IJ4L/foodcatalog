package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`

	MongoDBHost     string `mapstructure:"DB_HOST"`
	MongoDBPort     string `mapstructure:"DB_PORT"`
	MongoDBUser     string `mapstructure:"DB_USER"`
	MongoDBPassword string `mapstructure:"DB_PASSWORD"`
	MongoDB         string `mapstructure:"DB_NAME"`

	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
