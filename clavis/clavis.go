// Copyright (c) 2021 Satvik Reddy
package clavis

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenType is a specific type of token, eg. access or refresh
type TokenType string

const (
	// AccessToken is used for access of individual routes
	AccessToken TokenType = "access"
	// RefreshToken is used to refresh the access token
	RefreshToken TokenType = "refresh"
)

// TokenClaims is all the data encoded within a token
type TokenClaims struct {
	ID   int64     `json:"id"`
	Type TokenType `json:"type"`
	jwt.StandardClaims
}

// CreateClaims will construct a claims object
func CreateClaims(id int64, tokenType TokenType, expiresIn int64) TokenClaims {
	return TokenClaims{
		ID:   id,
		Type: tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: getExpTime(expiresIn),
		},
	}
}

func getExpTime(expiresIn int64) int64 {
	return time.Now().Unix() + expiresIn
}

// GenerateToken generates a new JWT token
func GenerateToken(tokenType TokenType, claims TokenClaims, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)
	return signedToken, err
}

// VerifyToken verifies a jwt token and returns the claims
func VerifyToken(signedString string, tokenType TokenType, key []byte) (TokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedString,
		(*TokenClaims)(nil),
		getKeyFunc(tokenType, key),
	)
	if err != nil {
		return TokenClaims{}, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return *claims, nil
	}
	return TokenClaims{}, errors.New("invalid token")
}

func getKeyFunc(tokenType TokenType, key []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if claims, ok := token.Claims.(*TokenClaims); ok {
			if claims.Type == tokenType {
				return key, nil
			}
			return nil, errors.New("invalid token type")
		}
		return nil, errors.New("invalid claims")
	}
}
