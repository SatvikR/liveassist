// Copyright (c) 2021 Satvik Reddy
package db

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

// Connect creates a go-pg connection.
func Connect(addr string, user string, password string, name string) {
	db = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: name,
	})
}

// Close will disconnect from the database
func Close() error {
	return db.Close()
}

// Setup will create schemas
func Setup() error {
	// return loadSchema()
	if err := Healthcheck(); err != nil {
		return err
	}
	log.Println("Connected to database")

	if err := loadSchema(); err != nil {
		return err
	}
	log.Println("Created tables")
	return nil
}

// Healthcheck pings the database to make sure it is connected properly.
func Healthcheck() error {
	err := db.Ping(context.Background())
	return err
}
