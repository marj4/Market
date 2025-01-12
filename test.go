package main

import (
	"Market/backend"
	"Market/backend/db"
	"Market/config"
	error2 "Market/error"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	DB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	fmt.Println("Connect to DB is succesful")

	login := "dada"
	password := "dada"
	email := "dada"

	_, hashPassword, err := hash(password)
	if err != nil {
		log.Fatal(err)
	}

	user := backend.User{
		Login:    login,
		Password: hashPassword,
		Email:    email,
	}

	if err := db.AddUser(DB, user); err != nil {
		log.Fatal(err)
	}
}

func hash(password string) ([]byte, string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, "", error2.Wrap("Cant hash password", err)
	}
	return hash, string(hash), nil
}
