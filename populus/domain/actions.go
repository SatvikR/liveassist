// Copyright (c) 2021 Satvik Reddy
package domain

import (
	"errors"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/populus/config"
	"github.com/SatvikR/liveassist/populus/db"
)

var (
	ErrHashFailed     error = errors.New("unable to hash password")
	ErrUserExists     error = errors.New("user already exists")
	ErrTokenGenFailed error = errors.New("unable to generate tokens")
)

// Signup creates a user and saves it. Can return ErrHashFailed, ErrUserExists,
// or ErrTokenGenFailed errors. If no errors, returns the access token, and refresh token
func Signup(username string, password string, email string) (string, string, error) {
	hashedpw, err := HashPW(password)
	if err != nil {
		return "", "", ErrHashFailed
	}
	id, err := db.CreateUser(username, hashedpw, email)
	if err != nil {
		return "", "", ErrUserExists
	}

	accTok, refTok, err := generateTokens(id)
	if err != nil {
		return "", "", ErrTokenGenFailed
	}

	return accTok, refTok, nil
}

func Login() {

}

func generateTokens(id int64) (string, string, error) {
	accTok, err := clavis.GenerateToken(
		clavis.AccessToken,
		clavis.CreateClaims(
			id,
			clavis.AccessToken,
			clavis.AccessTokenDuration,
		),
		config.AccessTokenKey,
	)
	if err != nil {
		return "", "", err
	}
	refTok, err := clavis.GenerateToken(
		clavis.RefreshToken,
		clavis.CreateClaims(
			id,
			clavis.RefreshToken,
			clavis.RefreshTokenDuration,
		),
		config.RefreshTokenKey,
	)
	if err != nil {
		return "", "", err
	}
	return accTok, refTok, nil
}

func Logout() {

}
