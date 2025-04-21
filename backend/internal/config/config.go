package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, relying on system environments")
	}

	viper.AutomaticEnv()
	viper.SetDefault("DB_URL", "DB_URL=postgres//test:testing@localhost:5433/postgres")
}

func Get(key string) string {
	return viper.GetString(key)
}
