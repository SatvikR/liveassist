// Copyright (c) 2021 Satvik Reddy
package messaging

import (
	"encoding/json"
	"log"

	"github.com/SatvikR/liveassist/nuntius/db"
	"github.com/SatvikR/liveassist/omnis/mq"
)

func listen() error {
	msgs, err := ch.Consume(
		queue.Name,
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
		for delivery := range msgs {
			var data mq.UserMessage
			err := json.Unmarshal(delivery.Body, &data)
			if err != nil {
				log.Printf("Invalid message recieved: %s", err.Error())
				continue
			}
			go handleMsg(data)
		}
	}()
	return nil
}

func handleMsg(data mq.UserMessage) error {
	switch data.Event {
	case mq.UserCreated:
		return db.SaveUser(data.Username, data.ID)
	default:
		return mq.ErrInvalidEvent
	}
}
