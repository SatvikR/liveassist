// Copyright (c) 2021 Satvik Reddy
package db

import (
	"context"
	"log"
	"time"

	"github.com/SatvikR/liveassist/nuntius/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var (
	client   *mongo.Client
	messages *mongo.Collection
)

const (
	messagesCollection string = "messages"
)

// Setup will connect to the mongodb database and setup everything.
func Setup() error {
	// Create client
	cs, err := connstring.ParseAndValidate(config.DBUri)
	if err != nil {
		return err
	}
	_db, err := mongo.NewClient(options.Client().ApplyURI(cs.String()))
	if err != nil {
		return err
	}
	client = _db
	// Connect client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Healthcheck
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}
	// Setup Collections
	messages = client.Database(cs.Database).Collection(messagesCollection)

	// Indexes
	name, err := messages.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "channelId", Value: 1},
				{Key: "createdAt", Value: 1},
			},
		},
	)
	if err != nil {
		return err
	}
	log.Printf("Created index %s", name)
	return nil
}

// Close disconnects from mongodb
func Close() {
	client.Disconnect(context.Background())
}
