// Copyright (c) 2021 Satvik Reddy
package http

import (
	"net/http"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/omnis"
	"github.com/SatvikR/liveassist/verum/config"
	"github.com/SatvikR/liveassist/verum/domain"
	"github.com/gin-gonic/gin"
)

func refresh(c *gin.Context) {
	refreshToken, err := c.Cookie(clavis.RefreshTokenCookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "cannot find refreshToken cookie",
		})
		return
	}
	accTok, refTok, err := domain.VerifyAndRefresh(refreshToken)
	switch err {
	case domain.ErrInvalidToken:
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	case omnis.ErrTokenGenFailed:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	clavis.SetRefreshTokenCookie(c, refTok, config.Domain)
	c.JSON(http.StatusAccepted, gin.H{
		"accessToken": accTok,
	})
}
