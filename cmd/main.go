package main

import (
	"Market/backend/db"
	"Market/config"
	"fmt"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	DB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}

	defer DB.Close()

	fmt.Println("Connect to DB is succesful")
}
