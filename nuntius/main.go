// Copyright (c) 2021 Satvik Reddy
package main

import (
	"log"

	"github.com/SatvikR/liveassist/nuntius/config"
	"github.com/SatvikR/liveassist/nuntius/db"
)

func main() {
	config.LoadConfig(8084)

	if err := db.Setup(); err != nil {
		log.Fatalf("Unable connect to database: %s", err.Error())
	}
	log.Println("Connected to database")
	defer db.Close()
}
