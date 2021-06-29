// Copyright (c) 2021 Satvik Reddy
package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type message struct {
	data []byte
}

type client struct {
	chanId string
	send   chan *message
	hub    *hub
	conn   *websocket.Conn
}

type channel struct {
	clients map[*client]bool
}

type hub struct {
	channels   map[string]*channel
	register   chan *client
	unregister chan *client
	broadcast  chan *message
}

func newHub() *hub {
	return &hub{
		channels:   make(map[string]*channel),
		register:   make(chan *client),
		unregister: make(chan *client),
		broadcast:  make(chan *message),
	}
}

func newChannel() *channel {
	return &channel{
		clients: make(map[*client]bool),
	}
}

func newClient(chanId string, hub *hub, conn *websocket.Conn) *client {
	return &client{
		chanId: chanId,
		send:   make(chan *message),
		hub:    hub,
		conn:   conn,
	}
}

func (h *hub) start() {
	for {
		select {
		case client := <-h.register:
			// If a channel already exists with the clients chanId, add the user to that channel
			if ch, ok := h.channels[client.chanId]; ok {
				ch.clients[client] = true
				break
			}
			// otherwise, add a new channel to the channels map and add the client to that
			h.channels[client.chanId] = newChannel()
			h.channels[client.chanId].clients[client] = true
		case client := <-h.unregister:
			// Remove client from channel
			delete(h.channels[client.chanId].clients, client)
			// Delete channel if it has no clients
			if len(h.channels[client.chanId].clients) < 1 {
				delete(h.channels, client.chanId)
			}
		}
	}
}

func (c *client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		log.Printf("Recieved message: %s", string(message))
	}
}

func serveWs(hub *hub, w http.ResponseWriter, r *http.Request, chanId string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if gin.Mode() == gin.DebugMode {
			log.Printf("Failed to connect client: %s", err.Error())
		}
		return
	}

	client := newClient(chanId, hub, conn)
	client.hub.register <- client

	go client.readPump()
}
