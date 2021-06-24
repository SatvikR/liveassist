// Copyright (c) 2021 Satvik Reddy
package config

import (
	"os"

	"github.com/SatvikR/liveassist/omnis"
	_ "github.com/joho/godotenv/autoload"
)

var (
	AccessTokenKey  []byte
	RefreshTokenKey []byte
	Port            int
	Domain          string
)

// LoadConfig loads the config from the environment
func LoadConfig(port int) {
	AccessTokenKey = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	RefreshTokenKey = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
	Port = omnis.GetPort(os.Getenv("PORT"), port)
	Domain = os.Getenv("DOMAIN")
}
