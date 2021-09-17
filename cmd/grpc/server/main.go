package main

import (
	"context"
	"log"
	"net"

	impl "github.com/gabrielfvale/klever-grpc/internal/grpc/impl"
	pb "github.com/gabrielfvale/klever-grpc/internal/proto-files"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	db       *mongo.Client
	cryptodb *mongo.Collection
	mongoCtx context.Context
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}

	s := grpc.NewServer()
	serviceServer := impl.NewCryptoServiceServer()
	pb.RegisterCryptoServiceServer(s, serviceServer)
	log.Printf("Listening on %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
