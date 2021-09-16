package main

import (
	"context"
	"log"
	"net"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto-files"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	db       *mongo.Client
	cryptodb *mongo.Collection
	mongoCtx context.Context
)

type CryptoServiceServer struct {
	*pb.UnimplementedCryptoServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCryptoServiceServer(s, &CryptoServiceServer{})
	log.Printf("Listening on %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
