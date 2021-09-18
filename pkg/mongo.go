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

const (
	DB         = "klever"
	COLLECTION = "crypto"
)

// GetMongoClient uses a singleton pattern to create a client
// connection to MongoDB, returning a pointer to mongo.Client
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		log.Printf("Connecting to MongoDB")
		opts := options.Client().ApplyURI(GetEnvVar("MONGO_URI"))

		client, err := mongo.Connect(mongoCtx, opts)
		if err != nil {
			mongoError = err
			return
		}

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

// GetMongoCollection returns the default Mongo collection
func GetMongoCollection() (*mongo.Collection, error) {
	// Get singleton client
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	// Return collection from const values
	collection := client.Database(DB).Collection(COLLECTION)
	return collection, nil
}
