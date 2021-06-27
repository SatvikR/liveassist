// Copyright (c) 2021 Satvik Reddy
package mq

import (
	"errors"

	"github.com/streadway/amqp"
)

const (
	AmnisQName string = "amnis"
)

// populus events
const (
	UserCreated string = "USER_CREATED"
)

var (
	ErrInvalidEvent error = errors.New("invalid event")
)

type UserMessage struct {
	Event    string `json:"event"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Connect connects to the amql address, and returns a connection and a channel.
// Make sure to defer closing both of those.
func Connect(addr string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	return conn, ch, nil
}

// GetQueue returns a durable queue with the name and channel given
func GetQueue(name string, ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
}
