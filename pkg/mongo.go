package pkg

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variables for singleton pattern
var (
	mongoClient *mongo.Client
	mongoCtx    = context.TODO()
	mongoError  error
	mongoOnce   sync.Once
)

// InitMongoClient uses a singleton pattern to create a client
// connection to MongoDB, returning a pointer to mongo.Client
func InitMongoClient(uri string) (*mongo.Client, error) {
	mongoOnce.Do(func() {
		log.Printf("Connecting to MongoDB")
		opts := options.Client().ApplyURI(uri)

		client, err := mongo.Connect(mongoCtx, opts)
		if err != nil {
			mongoError = err
			return
		}
		defer client.Disconnect(mongoCtx)

		err = client.Ping(mongoCtx, nil)
		if err != nil {
			mongoError = err
			return
		}
		mongoClient = client
		log.Print("Connected to Mongo successfully")
	})
	return mongoClient, mongoError
}
