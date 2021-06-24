// Copyright (c) 2021 Satvik Reddy
package http

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartServer will start an http server. Routes are:
// POST /signup, POST /login, DELETE /logout
func StartServer(port int, origins []string) {
	r := gin.Default()

	r.POST("/signup", signup)
	r.POST("/login", login)
	r.DELETE("/logout", logout)

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
