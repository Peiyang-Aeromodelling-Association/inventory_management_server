package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config interface {
	load(path string) error
}

type AppConfig struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
}

type SecretConfig struct {
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
}

// Exposed unified method for loading config
func LoadConfig(config Config, path string) error {
	return config.load(path)
}

// methods for loading AppConfig
func (config *AppConfig) load(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read config: ", err)
		return err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal("cannot unmarshal config: ", err)
		return err
	}

	return nil
}

// methods for loading SecretConfig
func (config *SecretConfig) load(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("secret")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read config: ", err)
		return err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal("cannot unmarshal config: ", err)
		return err
	}

	return nil
}
