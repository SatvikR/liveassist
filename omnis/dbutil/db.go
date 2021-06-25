// Copyright (c) 2021 Satvik Reddy
package dbutil

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// Connect creates a go-pg connection.
func Connect(db **pg.DB, addr string, user string, password string, name string) {
	*db = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: name,
	})
}

// Healthcheck pings the database to make sure it is connected properly.
func Healthcheck(db *pg.DB) error {
	err := db.Ping(context.Background())
	return err
}
