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

	DB2 := db.Connect2(cfg.Redis_Server)


	defer DB.Close()

	fmt.Println("Connect to DB is succesful")

	srv := server.StartServer(DB,DB2)
	if err := srv.Router.Run(":8080"); err != nil { log.Fatal(err) }
	fmt.Println("Server is started successfully")


}


