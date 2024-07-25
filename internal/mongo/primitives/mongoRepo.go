package primitives

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	Client   *mongo.Client
	Ctx      context.Context
	Cancel   context.CancelFunc
	schedule *Schedule
}

func NewMongoConn() *MongoStore {
	clientOptions := options.Client().ApplyURI("mongodb://admin:Erunda228@mongo")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoStore{
		Client: client,
		Ctx:    ctx,
		Cancel: cancel,
	}
}

func (s *MongoStore) Close() {
	s.Cancel()
}

func (s *MongoStore) Schedule() *Schedule {
	if s.schedule == nil {
		s.schedule = &Schedule{
			Store: s,
		}
	}
	return s.schedule
}
