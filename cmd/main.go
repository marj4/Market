package main

import (
	"Market/config"
	"Market/pkg/db"
	"Market/pkg/server"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil { log.Fatal(err) }


	DB, err := db.Connect(cfg.DatabaseURL)
	if err != nil { log.Fatal(err) }
	defer DB.Close()

	DB2 := db.Connect2(cfg.Redis_Server)
	fmt.Println("Connect to DB's is succesful")

	//Start server
	srv := server.StartServer(DB,DB2)
	if err := srv.Router.Run(":8080"); err != nil { log.Fatal(err) }
	fmt.Println("Server is started successfully")

	fmt.Println("Connect to DB is succesful")
}


