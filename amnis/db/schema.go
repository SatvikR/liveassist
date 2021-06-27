// Copyright (c) 2021 Satvik Reddy
package db

import "github.com/go-pg/pg/v10/orm"

type Channel struct {
	ID       string `pg:"type:uuid,default:uuid_generate_v4()"`
	OwnerID  int
	Name     string
	Keywords string

	Owner User `pg:"rel:has-one"`
}

type User struct {
	ID       int
	Username string
}

func loadSchema() error {
	if _, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`); err != nil {
		return err
	}

	for _, m := range []interface{}{(*Channel)(nil), (*User)(nil)} {
		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
