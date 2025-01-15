package main

import (
	"Market/backend/db"
	"Market/config"
	"fmt"
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

	password, err := db.GetUser(DB, "alina77")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(password)
}
