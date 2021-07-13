// Copyright (c) 2021 Satvik Reddy
package ws

import (
	"fmt"
	"net/http"

	"github.com/SatvikR/liveassist/clavis"
	"github.com/SatvikR/liveassist/nuntius/config"
	"github.com/SatvikR/liveassist/nuntius/domain"
	"github.com/SatvikR/liveassist/omnis"
	"github.com/gin-gonic/gin"
)

// StartServer starts the nuntius service
// Routes: Websocket (GET) /ws
func StartServer(port int) {
	hub := newHub()
	go hub.start()
	r := gin.Default()
	r.Use(omnis.GetCors())
	a := r.Use(clavis.JWTAuthURLMiddleware(config.AccessTokenKey))
	a.GET("/messages/ws", func(c *gin.Context) {
		chanId := c.Query("channel")
		if chanId != "" && domain.ChannelExists(chanId) {
			userId := c.GetInt64("uid")
			serveWs(hub, c.Writer, c.Request, chanId, int(userId))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid channel id",
			})
		}
	})

	r.Run(fmt.Sprintf(":%d", port))
}
