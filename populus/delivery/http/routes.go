// Copyright (c) 2021 Satvik Reddy
package http

import (
	"net/http"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/omnis"
	"github.com/SatvikR/liveassist/populus/config"
	"github.com/SatvikR/liveassist/populus/domain"
	"github.com/gin-gonic/gin"
)

type signupBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func signup(c *gin.Context) {
	var body signupBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": omnis.ErrInvalidBody.Error(),
		})
		return
	}

	accTok, refTok, err := domain.Signup(
		body.Username,
		body.Password,
		body.Email,
	)
	if err != nil {
		switch err {
		case domain.ErrHashFailed:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		case domain.ErrUserExists:
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		case omnis.ErrTokenGenFailed:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		return
	}

	clavis.SetRefreshTokenCookie(c, refTok, config.Domain)

	c.JSON(http.StatusCreated, gin.H{
		"accessToken": accTok,
	})
}

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c *gin.Context) {
	var body loginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": omnis.ErrInvalidBody.Error(),
		})
		return
	}

	accTok, refTok, err := domain.Login(body.Username, body.Password)
	if err != nil {
		switch err {
		case domain.ErrWrongPassword:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		case omnis.ErrTokenGenFailed:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		case domain.ErrUserNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		return
	}

	clavis.SetRefreshTokenCookie(c, refTok, config.Domain)
	c.JSON(http.StatusAccepted, gin.H{
		"accessToken": accTok,
	})
}

func logout(c *gin.Context) {
	clavis.SetRefreshTokenCookie(c, "", config.Domain)
	c.Status(http.StatusOK)
}
