// Copyright (c) 2021 Satvik Reddy
package ws

import (
	"fmt"
	"net/http"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/nuntius/config"
	"github.com/gin-gonic/gin"
)

// StartServer starts the nuntius service
// Routes: Websocket (GET) /ws
func StartServer(port int) {
	hub := newHub()
	go hub.start()
	r := gin.Default()
	a := r.Use(clavis.JWTAuthMiddleware(config.AccessTokenKey))
	a.GET("/ws", func(c *gin.Context) {
		chanId := c.Query("channel")
		if chanId != "" {
			serveWs(hub, c.Writer, c.Request, chanId)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "missing channel id",
			})
		}
	})

	r.Run(fmt.Sprintf(":%d", port))
}
