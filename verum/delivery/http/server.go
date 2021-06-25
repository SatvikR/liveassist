// Copyright (c) 2021 Satvik Reddy
package http

import (
	"fmt"

	"github.com/SatvikR/liveassist/omnis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartServer will start an http server. Routes are:
// PUT /refresh
func StartServer(port int) {
	r := gin.Default()

	g := r.Group("/api/tokens")

	g.PUT("/refresh", refresh)

	origins := omnis.GetDomain()

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
