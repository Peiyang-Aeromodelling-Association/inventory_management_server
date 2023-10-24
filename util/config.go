package util

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBHost              string        `mapstructure:"DB_HOST"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	PostgresPassword    string        `mapstructure:"POSTGRES_PASSWORD"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AdminUsername       string        `mapstructure:"ADMIN_USERNAME"`
	AdminPassword       string        `mapstructure:"ADMIN_PASSWORD"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// methods for loading config
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read config: ", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("cannot unmarshal config: ", err)
		return
	}

	return
}
