// Copyright (c) 2021 Satvik Reddy
package domain

import (
	"log"

	"github.com/SatvikR/liveassist/nuntius/db"
	"github.com/SatvikR/liveassist/omnis"
)

// SaveMessage saves a message
func SaveMessage(message, channelId string, userId int) (db.Message, error) {
	msgData, err := db.CreateMessage(message, channelId, userId)
	if err != nil {
		log.Printf("error: %v", err)
		return db.Message{}, omnis.ErrCouldNotCreate
	}
	return msgData, nil
}

// ChannelExists returns whether or not a channel exists
// You are probably wondering: "Oh really? I never would've guessed!".
// And to that I say: "yes."
func ChannelExists(chanId string) bool {
	_, err := db.FindChannel(chanId)
	return err == nil
}

// LoadMessages loads the initial messages for a channel. This is used when a
// client connects.
// TODO: pagination
func LoadMessages(chanId string) ([]db.Message, error) {
	messages, err := db.FindInChannel(chanId)
	if err != nil {
		return []db.Message{}, err
	}

	return messages, nil
}
