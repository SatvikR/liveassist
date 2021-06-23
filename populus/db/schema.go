// Copyright (c) 2021 Satvik Reddy
package db

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
)

type user struct {
	ID        int64
	Username  string `pg:",unique"`
	Password  string
	Email     string    `pg:",unique"`
	CreatedAt time.Time `pg:"default:now()"`
}

func loadSchema() error {
	err := db.Model((*user)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	return err
}
