// Copyright (c) 2021 Satvik Reddy
package main

import (
	"log"

	"github.com/SatvikR/liveassist/amnis/config"
	"github.com/SatvikR/liveassist/amnis/db"
	"github.com/SatvikR/liveassist/amnis/delivery/http"
	"github.com/SatvikR/liveassist/amnis/messaging"
)

func main() {
	config.LoadConfig(8082)

	if err := db.Setup(); err != nil {
		log.Fatalf("Unable to setup database: %s", err.Error())
	}
	defer db.Close()

	if err := messaging.Setup(); err != nil {
		log.Fatalf("Unable to setup messaging: %s", err.Error())
	}

	http.StartServer(config.Port)
}
