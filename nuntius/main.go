// Copyright (c) 2021 Satvik Reddy
package main

import (
	"log"

	"github.com/SatvikR/liveassist/nuntius/config"
	"github.com/SatvikR/liveassist/nuntius/db"
	"github.com/SatvikR/liveassist/nuntius/delivery/ws"
	"github.com/SatvikR/liveassist/nuntius/messaging"
)

func main() {
	config.LoadConfig(8084)

	if err := db.Setup(); err != nil {
		log.Fatalf("Unable connect to database: %s", err.Error())
	}
	log.Println("Connected to database")
	defer db.Close()

	if err := messaging.Setup(); err != nil {
		log.Fatalf("Unable to initialize rabbit mq connection")
	}

	ws.StartServer(config.Port)
}
