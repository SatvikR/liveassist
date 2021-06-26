// Copyright (c) 2021 Satvik Reddy
package domain

import (
	"errors"
	"strings"

	"github.com/SatvikR/liveassist/amnis/db"
	"github.com/SatvikR/liveassist/omnis"
)

var (
	ErrChannelDoesNotExist error = errors.New("channel does not exist")
	ErrCannotFindChannels  error = errors.New("cannot fetch channels")
)

type Channel struct {
	ID       string   `json:"id"`
	OwnerID  int      `json:"ownerId"`
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

// Create creates a channel and gives back the id
func Create(name string, ownerID int, keywords []string) (string, error) {
	// TODO make sure user exists
	// Convert keywords into a single string
	var kwstr strings.Builder
	for i, kw := range keywords {
		kwstr.WriteString(kw)
		if i != len(keywords)-1 {
			kwstr.WriteByte(' ')
		}
	}

	// Create Channel
	id, err := db.CreateChannel(name, ownerID, kwstr.String())
	// Error handling
	if err != nil {
		return "", omnis.ErrCouldNotCreate
	}
	return id, nil
}

// Delete deletes a channel from an id. Can return ErrChannelDoesNotExist.
func Delete(id string) error {
	err := db.DeleteChannel(id)
	if err != nil {
		return ErrChannelDoesNotExist
	}

	return nil
}

// TODO fix users embedding

// GetChannel gives back a single channel's data
func GetChannel(id string) (Channel, error) {
	c, err := db.FindChannel(id)
	if err != nil {
		return Channel{}, ErrChannelDoesNotExist
	}
	return constructChannel(c), nil
}

// GetChannels returns all of the channels
// TODO pagination
func GetChannels() ([]Channel, error) {
	c, err := db.FindAllChannels()
	if err != nil {
		return nil, ErrCannotFindChannels
	}

	out := make([]Channel, len(c))
	for i, ch := range c {
		out[i] = constructChannel(ch)
	}
	return out, nil
}

func constructChannel(c *db.Channel) Channel {
	return Channel{
		ID:       c.ID,
		Name:     c.Name,
		OwnerID:  c.OwnerID,
		Keywords: strings.Split(c.Keywords, " "),
	}
}
