package main

import (
	"log"
	"net"

	impl "github.com/gabrielfvale/klever-grpc/internal/grpc/impl"
	pb "github.com/gabrielfvale/klever-grpc/internal/proto-files"
	"github.com/gabrielfvale/klever-grpc/pkg"
	"google.golang.org/grpc"
)

func main() {
	// Initialize .env
	err := pkg.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to MongoDB
	_, err = pkg.InitMongoClient(pkg.GetEnvVar("MONGO_URI"))
	if err != nil {
		log.Fatalf("Could not initialize Mongo client: %v", err)
	}

	// Create listen socket
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}

	// Create gRPC server
	server := grpc.NewServer()
	serviceServer := impl.NewCryptoServiceServer()
	pb.RegisterCryptoServiceServer(server, serviceServer)
	log.Printf("Listening on %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
