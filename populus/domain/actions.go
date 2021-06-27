// Copyright (c) 2021 Satvik Reddy
package domain

import (
	"errors"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/omnis"
	"github.com/SatvikR/liveassist/populus/config"
	"github.com/SatvikR/liveassist/populus/db"
	"github.com/SatvikR/liveassist/populus/messaging"
)

// Errors that the functions can return
var (
	ErrHashFailed           error = errors.New("unable to hash password")
	ErrUserExists           error = errors.New("user already exists")
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
	id, username, err := db.CreateUser(username, hashedpw, email)
	if err != nil {
		return "", "", ErrUserExists
	}

	accTok, refTok, err := clavis.GenerateTokenPair(id, config.AccessTokenKey, config.RefreshTokenKey)
	if err != nil {
		return "", "", omnis.ErrTokenGenFailed
	}

	messaging.DispatchUserData(int(id), username)

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

	accTok, refTok, err := clavis.GenerateTokenPair(user.ID, config.AccessTokenKey, config.RefreshTokenKey)
	if err != nil {
		return "", "", omnis.ErrTokenGenFailed
	}

	return accTok, refTok, err
}
