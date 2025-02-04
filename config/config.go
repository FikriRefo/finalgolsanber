package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

var ENV Config

func LoadConfig() {
	viper.SetConfigFile(".env") // Set langsung ke file .env
	viper.SetConfigType("env")  // Format file
	viper.AutomaticEnv()        // Membaca dari environment jika ada

	// Membaca konfigurasi dari file .env
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading .env file: %s", err))
	}

	// Unmarshal ke struct
	if err := viper.Unmarshal(&ENV); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %s", err))
	}
}
