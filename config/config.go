package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT         string
	DATABASE_URL string
}

var ENV Config

func LoadConfig() {
	viper.AutomaticEnv() // Membaca dari environment

	ENV.PORT = viper.GetString("PORT")
	ENV.DATABASE_URL = viper.GetString("DATABASE_URL")

	if ENV.DATABASE_URL == "" {
		log.Fatal("DATABASE_URL is not set in environment variables")
	}

	fmt.Println("Config loaded successfully")
}
