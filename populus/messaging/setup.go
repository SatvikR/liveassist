// Copyright (c) 2021 Satvik Reddy
package messaging

import (
	"log"

	"github.com/SatvikR/liveassist/omnis/mq"
	"github.com/SatvikR/liveassist/populus/config"
	"github.com/streadway/amqp"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
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
	_, err = mq.GetFanoutExchange(mq.PopulusExchange, ch)
	if err != nil {
		return err
	}
	return nil
}

// Close will disconnect from rabbitmq.
func Close() {
	conn.Close()
	ch.Close()
}
