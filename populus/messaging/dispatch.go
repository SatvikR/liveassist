// Copyright (c) 2021 Satvik Reddy
package messaging

import (
	"encoding/json"

	"github.com/SatvikR/liveassist/omnis/mq"
	"github.com/streadway/amqp"
)

// DispatchUserData sends user data to amnis via rabbit mq
func DispatchUserData(id int, username string) error {
	body, err := json.Marshal(&mq.UserMessage{
		Event:    mq.UserCreated,
		ID:       id,
		Username: username,
	})
	if err != nil {
		return err
	}

	err = ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}
