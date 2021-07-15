// Copyright (c) 2021 Satvik Reddy
package messaging

import (
	"encoding/json"
	"log"

	"github.com/SatvikR/liveassist/nuntius/db"
	"github.com/SatvikR/liveassist/omnis/mq"
	"github.com/streadway/amqp"
)

func listen() error {
	userMsgs, err := ch.Consume(
		usersQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	chMsgs, err := ch.Consume(
		chQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for delivery := range userMsgs {
			go func(delivery amqp.Delivery) {
				var data mq.UserMessage
				err := json.Unmarshal(delivery.Body, &data)
				if err != nil {
					log.Printf("Invalid message recieved: %s", err.Error())
					return
				}
				go handleUserMsg(data)
			}(delivery)
		}
	}()
	go func() {
		for delivery := range chMsgs {
			go func(delivery amqp.Delivery) {
				var data mq.ChannelMessage
				err := json.Unmarshal(delivery.Body, &data)
				if err != nil {
					log.Printf("Invalid message recieved: %s", err.Error())
					return
				}
				go handleChMsg(data)
			}(delivery)
		}
	}()
	return nil
}

func handleUserMsg(data mq.UserMessage) error {
	switch data.Event {
	case mq.UserCreated:
		return db.SaveUser(data.Username, data.ID)
	default:
		return mq.ErrInvalidEvent
	}
}

func handleChMsg(data mq.ChannelMessage) error {
	switch data.Event {
	case mq.ChannelCreated:
		return db.SaveChannel(data.ID)
	case mq.ChannelDeleted:
		if err := db.DeleteInChannel(data.ID); err != nil {
			log.Printf("could not delete messages: %v", err)
			return err
		}
		if err := db.DeleteChannel(data.ID); err != nil {
			log.Printf("could not delete channel: %v", err)
			return err
		}
		return nil
	default:
		return mq.ErrInvalidEvent
	}
}
