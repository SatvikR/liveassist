// Copyright (c) 2021 Satvik Reddy
package mq

import (
	"errors"

	"github.com/streadway/amqp"
)

const (
	PopulusExchange string = "populus"
	AmnisExchange   string = "amnis"
	AmnisUsersQueue string = "amnis_users_q"
)

// events
const (
	UserCreated    string = "USER_CREATED"
	ChannelCreated string = "CHANNEL_CREATED"
	ChannelDeleted string = "CHANNEL_DELETED"
)

var (
	ErrInvalidEvent error = errors.New("invalid event")
)

type UserMessage struct {
	Event    string `json:"event"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type ChannelMessage struct {
	Event string `json:"event"`
	ID    string `json:"id"`
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

func GetNonDurableQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
}

// GetFanoutExchange creates a fanout exchange with the given name and returns the name
func GetFanoutExchange(name string, ch *amqp.Channel) (string, error) {
	err := ch.ExchangeDeclare(
		name,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return "", err
	}
	return name, nil
}

// BindQueue binds a queue to an exchange
func BindQueue(qName, exchange string, ch *amqp.Channel) error {
	return ch.QueueBind(
		qName,
		"",
		exchange,
		false,
		nil,
	)
}
