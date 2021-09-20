package main

import (
	"context"
	"log"
	"time"

	cli "github.com/gabrielfvale/klever-grpc/cmd/cobra"
	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"google.golang.org/grpc"
)

var (
	client pb.CryptoServiceClient
	ctx    context.Context
)

const (
	address = "localhost:50051"
)

func main() {
	log.Printf("Starting Crypto client")

	// Set up a connection to the gRPC server (insecure and blocking)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not dial: %v", err)
	}
	defer conn.Close()
	client = pb.NewCryptoServiceClient(conn)

	// Create timeout context
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Execute cobra CLI
	cli.Execute()
}
