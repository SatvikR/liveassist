// Copyright (c) 2021 Satvik Reddy
package clavis

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/SatvikR/liveassist/omnis"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenType is a specific type of token, eg. access or refresh
type TokenType string

const (
	// AccessToken is used for access of individual routes
	AccessToken TokenType = "access"
	// RefreshToken is used to refresh the access token
	RefreshToken TokenType = "refresh"
	// RefreshTokenCookie is the key for the refresh token cookie
	RefreshTokenCookie string = "liveassist_rtok"
)

// Token duration
const (
	RefreshTokenDuration int64 = 7 * 24 * 60 * 60
	AccessTokenDuration  int64 = 15
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
		// (*TokenClaims)(nil),
		&TokenClaims{},
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

// SetRefreshTokenCookie sets a refresh token cookie
func SetRefreshTokenCookie(c *gin.Context, refToken string, domain string) {
	c.SetCookie(
		RefreshTokenCookie,
		refToken,
		int(RefreshTokenDuration),
		omnis.RefreshRoute,
		domain,
		false,
		true,
	)
}

// GenerateTokenPair generates an accessToken and a refresh token from a user id
func GenerateTokenPair(id int64, accessTokenKey, refreshTokenKey []byte) (string, string, error) {
	accTok, err := GenerateToken(
		AccessToken,
		CreateClaims(
			id,
			AccessToken,
			AccessTokenDuration,
		),
		accessTokenKey,
	)
	if err != nil {
		return "", "", err
	}
	refTok, err := GenerateToken(
		RefreshToken,
		CreateClaims(
			id,
			RefreshToken,
			RefreshTokenDuration,
		),
		refreshTokenKey,
	)
	if err != nil {
		return "", "", err
	}
	return accTok, refTok, nil
}

type bearerHeader struct {
	Authorization string `header:"Authorization"`
}

// JWTAuthMiddleware returns a bearer token based middleware
func JWTAuthMiddleware(accessKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var header bearerHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "no authorization header",
			})
			return
		}

		var signedToken string
		if bearer := strings.Split(header.Authorization, " "); len(bearer) == 2 {
			signedToken = bearer[1]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			return
		}

		claims, err := VerifyToken(signedToken, AccessToken, accessKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}
		if claims.Type != AccessToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		c.Set("claims", &claims)
		c.Next()
	}
}
