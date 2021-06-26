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
)

// Create creates a channel and gives back the id
func Create(name string, ownerID int, keywords []string) (string, error) {
	// TODO make sure user exists
	// Convert keywords into a single string
	var kwstr strings.Builder
	for _, kw := range keywords {
		kwstr.WriteString(kw)
		kwstr.WriteByte(' ')
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

func Channel() {}

func Channels() {}
