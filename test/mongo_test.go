package test

import (
	"log"
	"os"
	"testing"

	"github.com/gabrielfvale/klever-grpc/pkg"
)

var mongoTestUri string = "mongodb://localhost:27018/"

// TestMongoInvalid tests, wby setting the MONGO_URI variable, if
// GetMongoClient() returns a valid *mongo.Client
func TestMongoValid(t *testing.T) {
	err := os.Setenv("MONGO_URI", "mongodb://localhost:27018/")
	if err == nil {
		client, err := pkg.GetMongoClient()

		if client == nil && err != nil {
			log.Fatal(err)
		}
	}
}

// TestMongoInvalid tests, when the Environ is empty, if GetMongoClient()
// returns an invalid *mongo.Client
func TestMongoSingleton(t *testing.T) {
	client, err := pkg.GetMongoClient()

	if client == nil || err != nil {
		t.Error("Invalid Mongo singleton")
	}
}

// TestCollectionInvalid tests, when the Environ is empty adn there is no client,
// if GetMongoCollection() returns an invalid *mongo.Collection
func TestMongoCollection(t *testing.T) {
	collection, err := pkg.GetMongoCollection()

	if collection == nil || err != nil {
		t.Error("Invalid Mongo collection")
	}
}
