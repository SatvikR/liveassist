// Copyright (c) 2021 Satvik Reddy
package db

import "github.com/go-pg/pg/v10"

// CreateChannel creates a channel and returns the id
func CreateChannel(name string, ownerID int, keywords string) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}

	newChannel := &Channel{
		Name:     name,
		OwnerID:  ownerID,
		Keywords: keywords,
	}
	if _, err := tx.Model(newChannel).Insert(); err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}
	return newChannel.ID, nil
}

// DeleteChannel deletes a channel given a uuid
func DeleteChannel(id string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Model((*Channel)(nil)).
		Where("? = ?", pg.Ident("id"), id).
		Delete(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// FindChannel finds a channel by a uuid
func FindChannel(id string) (*Channel, error) {
	channel := new(Channel)
	err := db.Model(channel).
		Relation("Owner").
		Where("? = ?", pg.Ident("channel.id"), id).
		Select()
	if err != nil {
		return nil, err
	}
	return channel, nil
}

// FindAllChannels returns all the channels
func FindAllChannels() ([]*Channel, error) {
	var channels []*Channel
	// TODO pagination
	err := db.Model(&channels).
		Relation("Owner").
		Select()
	if err != nil {
		return nil, err
	}

	return channels, nil
}

// SaveUser takes user data replicated from populus and saves it
func SaveUser(id int, username string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	newUser := &User{
		ID:       id,
		Username: username,
	}

	if _, err := tx.Model(newUser).Insert(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
