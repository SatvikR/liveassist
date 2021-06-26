// Copyright (c) 2021 Satvik Reddy
package main

import (
	"log"

	"github.com/SatvikR/liveassist/amnis/config"
	"github.com/SatvikR/liveassist/amnis/db"
	"github.com/SatvikR/liveassist/amnis/delivery/http"
)

func main() {
	config.LoadConfig(8082)

	if err := db.Setup(); err != nil {
		log.Fatalf("Unable to setup database: %s", err.Error())
	}
	defer db.Close()

	http.StartServer(config.Port)
}
