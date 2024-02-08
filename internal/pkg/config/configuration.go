package config

import (
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	ApiUrl       string `mapstructure:"API_URL"`
	Environment  string `mapstructure:"ENVIRONMENT"`
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`
	Port         string `mapstructure:"PORT"`
	Secret       string `mapstructure:"SECRET"`
}

func LoadConfig(path string) {
	var config *Configuration
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	Config = config
}

func GetConfig() *Configuration {
	return Config
}
