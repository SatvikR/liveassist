// Copyright (c) 2021 Satvik Reddy
package omnis

import "errors"

// Errors
var (
	ErrInvalidBody error = errors.New("invalid request body")
)

const (
	// RefreshRoute is the route used to refresh the access tokens and
	// is also the only route where we send refresh token cookies
	RefreshRoute string = "/refresh"
)
