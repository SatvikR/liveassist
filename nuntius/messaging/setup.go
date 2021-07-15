// Copyright (c) 2021 Satvik Reddy
package messaging

import (
	"log"

	"github.com/SatvikR/liveassist/nuntius/config"
	"github.com/SatvikR/liveassist/omnis/mq"
	"github.com/streadway/amqp"
)

var (
	conn       *amqp.Connection
	ch         *amqp.Channel
	usersQueue amqp.Queue
	chQueue    amqp.Queue
)

// Setup initializes the rabbit mq connection
func Setup() error {
	_conn, _ch, err := mq.Connect(config.MQUrl)
	if err != nil {
		return err
	}
	log.Println("Connected to rabbit mq")
	conn = _conn
	ch = _ch
	if err != nil {
		return err
	}
	if _, err := mq.GetFanoutExchange(mq.PopulusExchange, ch); err != nil {
		return err
	}
	if _, err := mq.GetFanoutExchange(mq.AmnisExchange, ch); err != nil {
		return err
	}
	_usersQueue, err := mq.GetQueue(mq.NuntiusUsersQueue, ch)
	if err != nil {
		return err
	}
	usersQueue = _usersQueue
	if err := mq.BindQueue(usersQueue.Name, mq.PopulusExchange, ch); err != nil {
		return err
	}
	_chQueue, err := mq.GetQueue(mq.NuntiusChannelsQueue, ch)
	if err != nil {
		return err
	}
	chQueue = _chQueue
	if err := mq.BindQueue(chQueue.Name, mq.AmnisExchange, ch); err != nil {
		return err
	}
	err = listen()
	if err != nil {
		return err
	}
	log.Println("Listening for messages")
	return nil
}

// Close will disconnect from rabbitmq.
func Close() {
	conn.Close()
	ch.Close()
}
