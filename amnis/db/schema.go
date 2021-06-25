// Copyright (c) 2021 Satvik Reddy
package db

import "github.com/go-pg/pg/v10/orm"

type Channel struct {
	ID       string `pg:"type:uuid,default:uuid_generate_v4()"`
	OwnerID  int
	Name     string
	Keywords string
}

func loadSchema() error {
	if _, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`); err != nil {
		return err
	}
	err := db.Model((*Channel)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	return err
}
