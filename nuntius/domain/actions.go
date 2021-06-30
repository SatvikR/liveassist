// Copyright (c) 2021 Satvik Reddy
package domain

import (
	"github.com/SatvikR/liveassist/nuntius/db"
	"github.com/SatvikR/liveassist/omnis"
)

// SaveMessage saves a message
func SaveMessage(message, channelId string, userId int) (db.Message, error) {
	msgData, err := db.CreateMessage(message, channelId, userId)
	if err != nil {
		return db.Message{}, omnis.ErrCouldNotCreate
	}
	return msgData, nil
}
