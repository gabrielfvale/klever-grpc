package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gabrielfvale/klever-grpc/internal/grpc/impl"
	"github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/gabrielfvale/klever-grpc/pkg"
	"google.golang.org/grpc"
)

var (
	serviceServer *impl.CryptoServiceServer
)

func init() {
	err := os.Setenv("MONGO_URI", mongoTestUri)

	// Connect to MongoDB
	_, err = pkg.GetMongoClient()
	if err != nil {
		log.Fatalf("Could not connect to Mongo")
	}

	// Create gRPC server
	server := grpc.NewServer()
	serviceServer = impl.NewCryptoServiceServer()
	proto.RegisterCryptoServiceServer(server, serviceServer)
}

// TestCreate tests if a new crypto can be created, also testing if two records
// with the same "symbol" field can be added.
func TestCreate(t *testing.T) {

	// Delete from test database
	serviceServer.DeleteCrypto(context.Background(), &proto.DeleteReq{Symbol: "TST"})

	res, err := serviceServer.CreateCrypto(context.Background(), &proto.CreateReq{
		Crypto: &proto.Crypto{
			Symbol: "TST",
			Name:   "Test",
		},
	})

	if err != nil || res.GetCrypto().Symbol != "TST" {
		t.Error("Invalid insertion")
	}

	// Try to insert again
	res, err = serviceServer.CreateCrypto(context.Background(), &proto.CreateReq{
		Crypto: &proto.Crypto{
			Symbol: "TST",
			Name:   "Test",
		},
	})

	if err == nil {
		t.Error("Double insertion with same symbol")
	}
}

// TestRead tests reading an invalid crypto from the database.
func TestRead(t *testing.T) {
	_, err := serviceServer.ReadCrypto(context.Background(), &proto.ReadReq{Symbol: "invalid"})
	if err == nil {
		t.Error("Read an invalid crypto")
	}

	res, err := serviceServer.ReadCrypto(context.Background(), &proto.ReadReq{Symbol: "TST"})
	if err != nil || res.GetCrypto().Name != "Test" {
		t.Error("Could not read a crypto")
	}
}

// TestUpdate tests updating a previously created "TST" record.
func TestUpdate(t *testing.T) {
	res, err := serviceServer.UpdateCrypto(context.Background(), &proto.UpdateReq{
		Crypto: &proto.Crypto{
			Symbol: "TST",
			Name:   "Testing",
		},
	})

	if err != nil || res.GetCrypto().Name != "Testing" {
		t.Error("Could update a crypto")
	}
}

// TestUpvote tests upvoting a Crypto record.
func TestUpvote(t *testing.T) {
	// Read previous value
	previous, err := serviceServer.ReadCrypto(context.Background(), &proto.ReadReq{Symbol: "TST"})

	// Upvote crypto
	res, err := serviceServer.Upvote(context.Background(), &proto.VoteRequest{Symbol: "TST"})

	if err != nil || res.Symbol != "TST" {
		t.Error("Error upvoting crypto")
	}

	// Read new value
	new, err := serviceServer.ReadCrypto(context.Background(), &proto.ReadReq{Symbol: "TST"})

	if err != nil || previous.GetCrypto().Upvotes > new.GetCrypto().Upvotes {
		t.Error("Failed to upvote crypto", err)
	}
}

// TestDownvote tests downvoting a Crypto record.
func TestDownvote(t *testing.T) {
	// Read previous value
	previous, err := serviceServer.ReadCrypto(context.Background(), &proto.ReadReq{Symbol: "TST"})

	// Downvote crypto
	res, err := serviceServer.Downvote(context.Background(), &proto.VoteRequest{Symbol: "TST"})

	if err != nil || res.Symbol != "TST" {
		t.Error("Error upvoting crypto")
	}

	// Read new value
	new, err := serviceServer.ReadCrypto(context.Background(), &proto.ReadReq{Symbol: "TST"})

	if err != nil || previous.GetCrypto().Downvotes > new.GetCrypto().Downvotes {
		t.Error("Failed to upvote crypto", err)
	}
}

// TestDelete tests deleting a crypto record.
func TestDelete(t *testing.T) {
	res, err := serviceServer.DeleteCrypto(context.Background(), &proto.DeleteReq{Symbol: "invalid"})
	if err == nil || res.Ok {
		t.Error("Invalid deletion")
	}

	res, err = serviceServer.DeleteCrypto(context.Background(), &proto.DeleteReq{Symbol: "TST"})
	if err != nil || !res.Ok {
		t.Error("Could not delete crypto")
	}
}
