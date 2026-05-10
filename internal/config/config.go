package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB   Database
	SMTP SMTP
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type SMTP struct {
	Host     string
	Port     string
	User     string
	Password string
	From     string
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpFrom := os.Getenv("SMTP_FROM")

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
		SMTP: SMTP{
			Host:     smtpHost,
			Port:     smtpPort,
			User:     smtpUser,
			Password: smtpPassword,
			From:     smtpFrom,
		},
	}

	return &ctf, nil
}
