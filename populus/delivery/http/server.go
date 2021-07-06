// Copyright (c) 2021 Satvik Reddy
package http

import (
	"fmt"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/omnis"
	"github.com/SatvikR/liveassist/populus/config"
	"github.com/gin-gonic/gin"
)

// StartServer will start an http server. Routes are:
// POST /signup, POST /login, DELETE /logout
func StartServer(port int) {
	r := gin.Default()
	r.Use(omnis.GetCors())

	g := r.Group("/api/users")
	a := r.Group("/api/users").Use(clavis.JWTAuthMiddleware(config.AccessTokenKey))

	g.POST("/signup", signup)
	g.POST("/login", login)
	g.DELETE("/logout", logout)
	a.GET("/me", me)

	r.Run(fmt.Sprintf(":%d", port))
}
