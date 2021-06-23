// Copyright (c) 2021 Satvik Reddy
package main

import (
	"log"

	"github.com/SatvikR/liveassist/populus/config"
	"github.com/SatvikR/liveassist/populus/db"
)

func main() {
	Setup()
}

// Setup will create any neccessary connections and variables needed to
// start the service.
func Setup() {
	config.LoadConfig()

	db.Connect(
		config.DBAddr,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	if err := db.Healthcheck(); err != nil {
		log.Fatalf("Failed to connect to db: %s\n", err)
	}

	log.Println("Connected to database")

}
