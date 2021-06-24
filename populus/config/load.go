// Copyright (c) 2021 Satvik Reddy
package config

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DBAddr          string
	DBUser          string
	DBPassword      string
	DBName          string
	AccessTokenKey  []byte
	RefreshTokenKey []byte
	Domain          string
	PORT            int
)

// LoadConfig loads configuratin data from environment.
func LoadConfig(defaultPort int) {
	DBAddr = os.Getenv("DB_ADDR")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	AccessTokenKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	RefreshTokenKey = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
	Domain = os.Getenv("DOMAIN")
	PORT = getPort(os.Getenv("PORT"), defaultPort)
}

func getPort(portString string, defaultPort int) int {
	p, err := strconv.Atoi(portString)
	if err != nil {
		log.Print("PORT environment variable not found, using default")
		return defaultPort
	} else {
		return p
	}
}
