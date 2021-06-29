// Copyright (c) 2021 Satvik Reddy
package db

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	Text      string             `bson:"text" json:"text"`
	// TODO change userId to replicated user data
	UserID    int    `bson:"userId" json:"userId"`
	ChannelID string `bson:"channelId" json:"channelId"`
}

type User struct {
	Uid      int    `bson:"uid" json:"id,omitempty"`
	Username string `bson:"username" json:"username"`
}

// CreateMessage stores a message in the database
func CreateMessage(text string, chanId string, userId int) (string, error) {
	newMessage := bson.D{
		{Key: "text", Value: text},
		{Key: "channelId", Value: chanId},
		{Key: "createdAt", Value: time.Now()},
		{Key: "userId", Value: userId},
	}
	_id, err := messages.InsertOne(
		context.Background(),
		newMessage,
	)
	if err != nil {
		return "", err
	}
	id, ok := _id.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("unable to parse id")
	}

	return id.Hex(), nil
}

// FindInChannel finds all the messages in a channel ordered by date
// TODO pagination(?, maybe not)
func FindInChannel(chanId string) ([]Message, error) {
	ctx := context.Background()

	queryFilter := bson.D{{Key: "channelId", Value: chanId}}
	queryOptions := &options.FindOptions{}
	queryOptions.SetSort(bson.D{{Key: "createdAt", Value: 1}})

	cur, err := messages.Find(ctx, queryFilter, queryOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var messages []Message
	if err := cur.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// SaveUser saves replicated user data
func SaveUser(username string, uid int) error {
	newUser := bson.D{
		{Key: "uid", Value: uid},
		{Key: "username", Value: username},
	}

	_, err := users.InsertOne(
		context.Background(),
		newUser,
	)
	return err
}
