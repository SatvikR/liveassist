// Copyright (c) 2021 Satvik Reddy
package omnis

import (
	"errors"
	"log"
	"strconv"
)

// Errors
var (
	ErrInvalidBody    error = errors.New("invalid request body")
	ErrTokenGenFailed error = errors.New("unable to generate tokens")
)

const (
	// RefreshRoute is the route used to refresh the access tokens and
	// is also the only route where we send refresh token cookies
	RefreshRoute string = "/refresh"
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
