package config

import (
	error2 "Market/error"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load("File.env"); err != nil {
		return nil, error2.Wrap("Failed load .env file", err)
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE"),
		Port:        os.Getenv("SERVER_PORT"),
	}, nil

}
