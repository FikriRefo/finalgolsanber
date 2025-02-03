package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	// Gunakan ENV struct dengan field yang sesuai
	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		ENV.DB_HOST, ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_NAME, ENV.DB_PORT)

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %s", err))
	}

	DB = db
}
