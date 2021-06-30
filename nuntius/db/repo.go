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
	User      User               `bson:"user" json:"owner"`
	ChannelID string             `bson:"channelId" json:"channelId"`
}

type User struct {
	Uid      int    `bson:"uid" json:"id,omitempty"`
	Username string `bson:"username" json:"username"`
}

// CreateMessage stores a message in the database
func CreateMessage(text string, chanId string, userId int) (Message, error) {
	user, err := GetUser(userId)
	if err != nil {
		return Message{}, err
	}

	msgData := Message{
		CreatedAt: time.Now(),
		Text:      text,
		User:      user,
		ChannelID: chanId,
	}
	newMessage := bson.D{
		{Key: "text", Value: msgData.Text},
		{Key: "channelId", Value: msgData.ChannelID},
		{Key: "createdAt", Value: msgData.CreatedAt},
		{Key: "user", Value: bson.D{
			{Key: "uid", Value: msgData.User.Uid},
			{Key: "username", Value: msgData.User.Username},
		}},
	}
	_id, err := messages.InsertOne(
		context.Background(),
		newMessage,
	)
	if err != nil {
		return Message{}, err
	}
	id, ok := _id.InsertedID.(primitive.ObjectID)
	msgData.ID = id
	if !ok {
		return Message{}, errors.New("unable to parse id")
	}

	return msgData, nil
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

// GetUser finds a user and returns the user info
func GetUser(uid int) (User, error) {
	queryFilter := bson.D{{Key: "uid", Value: uid}}
	res := users.FindOne(
		context.Background(),
		queryFilter,
	)
	var user User
	err := res.Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, err
}
