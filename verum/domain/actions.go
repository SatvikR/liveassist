// Copyright (c) 2021 Satvik Reddy
package domain

import (
	"errors"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/omnis"
	"github.com/SatvikR/liveassist/verum/config"
)

// Errors that the actions can return
var (
	ErrInvalidToken error = errors.New("invalid refresh token")
)

// VerifyAndRefresh verifies a refresh token and sends back a new refresh token
// and access token if it is valid. Can return the ErrInvalidToken or ErrTokenGenFailed errors
func VerifyAndRefresh(signedToken string) (string, string, error) {
	claims, err := clavis.VerifyToken(
		signedToken,
		clavis.RefreshToken,
		config.RefreshTokenKey,
	)
	if err != nil {
		return "", "", ErrInvalidToken
	}

	accTok, refTok, err := clavis.GenerateTokenPair(claims.ID, config.AccessTokenKey, config.RefreshTokenKey)
	if err != nil {
		return "", "", omnis.ErrTokenGenFailed
	}

	return accTok, refTok, nil
}
