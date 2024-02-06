package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	Port     string `mapstructure:"PORT"`
	Secret   string `mapstructure:"SECRET"`
}

func LoadConfig(path string) {
	var config *Configuration
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshal config: %v", err)
	}

	Config = config
}

func GetConfig() *Configuration {
	return Config
}
