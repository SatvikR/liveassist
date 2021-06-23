// Copyright (c) 2021 Satvik Reddy
package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DBAddr     string
	DBUser     string
	DBPassword string
	DBName     string
)

// LoadConfig loads configuratin data from environment.
func LoadConfig() {
	DBAddr = os.Getenv("DB_ADDR")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
}
