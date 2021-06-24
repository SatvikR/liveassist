// Copyright (c) 2021 Satvik Reddy
package main

import (
	"log"

	"github.com/SatvikR/liveassist/populus/config"
	"github.com/SatvikR/liveassist/populus/db"
	"github.com/SatvikR/liveassist/populus/delivery/http"
)

func main() {
	Setup()
}

// Setup will create any neccessary connections and variables needed to
// start the service.
func Setup() {
	config.LoadConfig(8080)

	db.Connect(
		config.DBAddr,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
	defer db.Close()
	if err := db.Setup(); err != nil {
		log.Fatalf("Unable to setup database: %s", err.Error())
	}

	http.StartServer(config.Port)
}
