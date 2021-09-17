package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto-files"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the gRPC server (insecure and blocking)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not dial: %v", err)
	}
	defer conn.Close()

	// Create the client and context
	client := pb.NewCryptoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test the endpoints
	_, err = client.Upvote(ctx, &pb.VoteRequest{Symbol: "BTC"})
	_, err = client.Downvote(ctx, &pb.VoteRequest{Symbol: "BTC"})
	_, err = client.CreateCrypto(ctx, &pb.CreateReq{Crypto: &pb.Crypto{Id: 1, Name: "Bitcoin", Symbol: "BTC"}})
	_, err = client.ReadCrypto(ctx, &pb.ReadReq{Query: &pb.ReadReq_Id{Id: 1}})
	_, err = client.UpdateCrypto(ctx, &pb.UpdateReq{Crypto: &pb.Crypto{Id: 1, Name: "Bitcoin", Symbol: "BTC"}})
	_, err = client.DeleteCrypto(ctx, &pb.DeleteReq{Query: &pb.DeleteReq_Id{Id: 1}})

	stream, err := client.ListCrypto(ctx, &pb.Empty{})
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListCrypto(_) = _, %v", client, err)
		}
		log.Println(res.Symbol)
	}

	if err != nil {
		log.Fatalf("could not upvote: %v", err)
	}
}
