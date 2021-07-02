// Copyright (c) 2021 Satvik Reddy
package http

import (
	"fmt"

	"github.com/SatvikR/liveassist/amnis/config"
	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/omnis"
	"github.com/gin-gonic/gin"
)

// StartServer will start an http server. Routes are:
// POST /, DELETE /{id}, GET /, GET /{id}
func StartServer(port int) {
	r := gin.Default()

	r.Use(omnis.GetCors())
	g := r.Group("/api/channels")
	a := r.Group("/api/channels")

	a.Use(clavis.JWTAuthMiddleware(config.AccessTokenKey))

	a.POST("/", create)
	a.DELETE("/:id", delete)
	g.GET("/:id", channel)
	g.GET("/", channels)

	r.Run(fmt.Sprintf(":%d", port))
}
