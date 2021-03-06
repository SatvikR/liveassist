// Copyright (c) 2021 Satvik Reddy
package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/SatvikR/liveassist/nuntius/db"
	"github.com/SatvikR/liveassist/nuntius/domain"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type messageBody struct {
	Text     *string    `json:"text"`
	LoadMore *time.Time `json:"loadMore"`
}

type message struct {
	data   []byte
	chanId string
}

type client struct {
	chanId string
	send   chan *message
	hub    *hub
	conn   *websocket.Conn
	userId int
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

func newClient(chanId string, hub *hub, conn *websocket.Conn, userId int) *client {
	return &client{
		chanId: chanId,
		send:   make(chan *message),
		hub:    hub,
		conn:   conn,
		userId: userId,
	}
}

func newMessage(msgObj db.Message) (*message, error) {
	data, err := json.Marshal(&msgObj)
	if err != nil {
		return nil, err
	}

	return &message{
		data:   data,
		chanId: msgObj.ChannelID,
	}, nil
}

func rawMessage(body interface{}, chanId string) (*message, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return &message{
		data:   data,
		chanId: chanId,
	}, nil
}

func (h *hub) start() {
	for {
		select {
		case client := <-h.register:
			// send the client initial channel data
			messages, err := domain.LoadMessages(client.chanId, time.Now())
			if err != nil {
				client.conn.Close()
				break
			}
			message, err := rawMessage(messages, client.chanId)
			if err != nil {
				client.conn.Close()
				break
			}
			client.send <- message
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
		case message := <-h.broadcast:
			// Check if the channel is on this hub
			if _, ok := h.channels[message.chanId]; !ok {
				break
			}
			// Send the message to each of the clients in the channel
			for client := range h.channels[message.chanId].clients {
				client.send <- message
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
			break
		}
		var msgBody messageBody
		if err = json.Unmarshal(message, &msgBody); err == nil {
			if msgBody.Text != nil {
				msgObj, err := domain.SaveMessage(*msgBody.Text, c.chanId, c.userId)
				if err != nil {
					log.Printf("error: %v", err)
					continue
				}
				if pmessage, err := newMessage(msgObj); err == nil {
					c.hub.broadcast <- pmessage
				}
			}
			if msgBody.LoadMore != nil {
				messages, err := domain.LoadMessages(c.chanId, *msgBody.LoadMore)
				if err != nil {
					log.Printf("error: %v", err)
					continue
				}
				messageObj, err := rawMessage(messages, c.chanId)
				if err != nil {
					log.Printf("error: %v", err)
					continue
				}
				c.send <- messageObj
			}
		}
	}
}

func (c *client) writePump() {
	ticker := time.NewTicker(60 * time.Second * 9 / 10) // Almost 60 seconds

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message.data)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func serveWs(hub *hub, w http.ResponseWriter, r *http.Request, chanId string, userId int) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if gin.Mode() == gin.DebugMode {
			log.Printf("Failed to connect client: %s", err.Error())
		}
		return
	}

	client := newClient(chanId, hub, conn, userId)
	client.hub.register <- client

	go client.readPump()
	go client.writePump()
}
