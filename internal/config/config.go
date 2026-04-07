package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("Error loading .env file", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	if host == "" || port == "" || user == "" || password == "" || name == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}

	if sslMode == "" {
		sslMode = "disable"
	}

	ctf := Config{
		DB: Database{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Name:     name,
			SSLMode:  sslMode,
		},
	}

	return &ctf, nil
}
