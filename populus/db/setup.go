// Copyright (c) 2021 Satvik Reddy
package db

import (
	"context"

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

// Healthcheck pings the database to make sure it is connected properly.
func Healthcheck() error {
	err := db.Ping(context.Background())
	return err
}
