// Copyright (c) 2021 Satvik Reddy
package http

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartServer will start an http server. Routes are:
// POST /signup, POST /login, DELETE /logout
func StartServer(port int) {
	r := gin.Default()

	g := r.Group("/api/users")

	g.POST("/signup", signup)
	g.POST("/login", login)
	g.DELETE("/logout", logout)

	var origins []string

	if gin.Mode() == gin.ReleaseMode {
		// TODO change this url to the prod url
		origins = []string{"http://localhost:3000"}
	} else {
		origins = []string{"http://localhost:3000"}
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = origins
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
	}
	r.Use(cors.New(corsConfig))

	r.Run(fmt.Sprintf(":%d", port))
}
