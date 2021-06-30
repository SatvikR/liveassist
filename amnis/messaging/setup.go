// Copyright (c) 2021 Satvik Reddy
package messaging

import (
	"log"

	"github.com/SatvikR/liveassist/amnis/config"
	"github.com/SatvikR/liveassist/omnis/mq"
	"github.com/streadway/amqp"
)

var (
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
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
	if _, err := mq.GetFanoutExchange(mq.PopulusExchange, ch); err != nil {
		return err
	}
	_queue, err := mq.GetNonDurableQueue(ch)
	if err != nil {
		return err
	}

	queue = _queue
	if err := mq.BindQueue(queue.Name, mq.PopulusExchange, ch); err != nil {
		return err
	}
	log.Println("Created amnis queue")
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
