// Copyright (c) 2021 Satvik Reddy
package omnis

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Errors
var (
	ErrInvalidBody    error = errors.New("invalid request body")
	ErrTokenGenFailed error = errors.New("unable to generate tokens")
	ErrCouldNotCreate error = errors.New("unable to create object")
)

const (
	// RefreshRoute is the route used to refresh the access tokens and
	// is also the only route where we send refresh token cookies
	RefreshRoute string = "/api/tokens/refresh"
)

func GetPort(portString string, defaultPort int) int {
	p, err := strconv.Atoi(portString)
	if err != nil {
		log.Print("PORT environment variable not found, using default")
		return defaultPort
	} else {
		return p
	}
}

func GetDomain() []string {
	var origins []string
	if gin.Mode() == gin.ReleaseMode {
		origins = []string{"https://liveassist.satvikreddy.com"}
	} else {
		origins = []string{"http://localhost:3000"}
	}
	return origins
}

func GetCors() gin.HandlerFunc {
	origins := GetDomain()
	corsConfig := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	cors := cors.New(corsConfig)
	return cors
}
