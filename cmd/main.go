package main

import (
	"Market/backend/db"
	"Market/backend/server"
	"Market/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil { log.Fatal(err) }

	DB, err := db.Connect(cfg.DatabaseURL)
	if err != nil { log.Fatal(err) }

	defer DB.Close()

	fmt.Println("Connect to DB is succesful")

	srv := server.StartServer(DB)
	if err := srv.Router.Run(":8080"); err != nil { log.Fatal(err) }
	fmt.Println("Server is started successfully")


}


