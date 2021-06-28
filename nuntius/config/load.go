// Copyright (c) 2021 Satvik Reddy
package config

import (
	"os"

	"github.com/SatvikR/liveassist/omnis"
	_ "github.com/joho/godotenv/autoload"
)

var (
	DBName         string
	DBUri          string
	AccessTokenKey []byte
	Port           int
)

// LoadConfig will load configuration details form the environment
func LoadConfig(defaultPort int) {
	DBName = os.Getenv("DB_NAME")
	DBUri = os.Getenv("DB_URI")
	AccessTokenKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	Port = omnis.GetPort(os.Getenv("PORT"), defaultPort)
}
