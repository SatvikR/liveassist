// Copyright (c) 2021 Satvik Reddy
package domain

import (
	"errors"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/populus/config"
	"github.com/SatvikR/liveassist/populus/db"
)

// Errors that the functions can return
var (
	ErrHashFailed           error = errors.New("unable to hash password")
	ErrUserExists           error = errors.New("user already exists")
	ErrTokenGenFailed       error = errors.New("unable to generate tokens")
	ErrWrongPassword        error = errors.New("incorrect password")
	ErrUserNotFound         error = errors.New("user does not exist")
	ErrPWVerificationFailed error = errors.New("could not verify passowrd")
)

// Signup creates a user and saves it. Can return the ErrHashFailed, ErrUserExists,
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

// Login verifies a user's credentials. Can return the ErrWrongPassword, ErrTokenGenFailed,
// or ErrUserNotFound errors. If no errors exist, returns the access token and refresh token
func Login(username string, password string) (string, string, error) {
	user, err := db.FindUserByUsername(username)
	if err != nil {
		return "", "", ErrUserNotFound
	}

	ok := VerifyPW(password, user.Password)
	if !ok {
		return "", "", ErrWrongPassword
	}

	accTok, refTok, err := generateTokens(user.ID)
	if err != nil {
		return "", "", ErrTokenGenFailed
	}

	return accTok, refTok, err
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
