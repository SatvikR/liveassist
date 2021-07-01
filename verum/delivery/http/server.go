// Copyright (c) 2021 Satvik Reddy
package http

import (
	"fmt"

	"github.com/SatvikR/liveassist/omnis"
	"github.com/gin-gonic/gin"
)

// StartServer will start an http server. Routes are:
// PUT /refresh
func StartServer(port int) {
	r := gin.Default()

	r.Use(omnis.GetCors())
	g := r.Group("/api/tokens")

	g.PUT("/refresh", refresh)

	r.Run(fmt.Sprintf(":%d", port))
}
