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
	a.GET("/messages/ws", func(c *gin.Context) {
		// TODO make sure channel exists
		chanId := c.Query("channel")
		if chanId != "" {
			userId := c.GetInt64("uid")
			serveWs(hub, c.Writer, c.Request, chanId, int(userId))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "missing channel id",
			})
		}
	})

	r.Run(fmt.Sprintf(":%d", port))
}
