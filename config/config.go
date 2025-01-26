package config

import (
	error2 "Market/error"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL  string //Адрес ДБ
	Port         string //Порт БД
	Email        string //Почта откуда будут отправляться коды верификации
	App_Password string //Пароль приложения для подключения к SMTP серверу(код верификации)
	Redis_Server string //Адрес сервера-Redis, в котором будут храниться временные данные
}

func LoadConfig() (*Config, error) {

	if err := godotenv.Load("File.env"); err != nil {
		return nil, error2.Wrap("Failed load .env file", err)
	}

	return &Config{
		DatabaseURL:  os.Getenv("DATABASE"),
		Port:         os.Getenv("SERVER_PORT"),
		Email:        os.Getenv("EMAIL"),
		App_Password: os.Getenv("APP_PASSWORD"),
		Redis_Server: os.Getenv("REDIS_SERVER"),
	}, nil

}
