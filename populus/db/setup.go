// Copyright (c) 2021 Satvik Reddy
package db

import (
	"log"

	"github.com/SatvikR/liveassist/omnis/dbutil"
	"github.com/SatvikR/liveassist/populus/config"
	"github.com/go-pg/pg/v10"
)

var db *pg.DB

// Setup will create schemas
func Setup() error {
	dbutil.Connect(&db, config.DBAddr, config.DBUser, config.DBPassword, config.DBName)

	if err := dbutil.Healthcheck(db); err != nil {
		return err
	}
	log.Println("Connected to database")

	if err := loadSchema(); err != nil {
		return err
	}
	log.Println("Created tables")
	return nil
}

// Close will close the database
func Close() error {
	return db.Close()
}
