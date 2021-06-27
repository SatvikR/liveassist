// Copyright (c) 2021 Satvik Reddy
package config

import (
	"os"

	"github.com/SatvikR/liveassist/omnis"
	_ "github.com/joho/godotenv/autoload"
)

var (
	DBAddr         string
	DBUser         string
	DBPassword     string
	DBName         string
	AccessTokenKey []byte
	MQUrl          string
	Port           int
)

// LoadConfig will load configuration details form the environment
func LoadConfig(defaultPort int) {
	DBAddr = os.Getenv("DB_ADDR")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	AccessTokenKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	MQUrl = os.Getenv("MQ_URL")
	Port = omnis.GetPort(os.Getenv("PORT"), defaultPort)
}
