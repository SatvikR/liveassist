// Copyright (c) 2021 Satvik Reddy
package messaging

import (
	"encoding/json"

	"github.com/SatvikR/liveassist/omnis/mq"
	"github.com/streadway/amqp"
)

// DispatchChannelData dispatches the channel id to a rabbit mq fanout
func DispatchChannelData(chanId string) error {
	body, err := json.Marshal(&mq.ChannelMessage{
		Event: mq.ChannelCreated,
		ID:    chanId,
	})
	if err != nil {
		return err
	}

	err = ch.Publish(
		mq.AmnisExchange,
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
