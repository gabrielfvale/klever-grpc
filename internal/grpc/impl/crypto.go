package impl

import (
	"context"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto-files"
)

// CryptoServiceServer is a implementation of CryptoService provided by gRPC
type CryptoServiceServer struct {
	*pb.UnimplementedCryptoServiceServer
}

// NewCryptoServiceServer returns a pointer to a CryptoServiceServer
func NewCryptoServiceServer() *CryptoServiceServer {
	return &CryptoServiceServer{}
}

// Upvote takes a VoteRequest and updates the "upvotes" field on a given crypto
// returning a VoteResponse if successful.
func (s *CryptoServiceServer) Upvote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteResponse, error) {
	log.Printf("Received Upvote request for %v", in.GetSymbol())
	return &pb.VoteResponse{}, nil
}

// Downvote takes a VoteRequest and updates the "downvotes" field on a given crypto
// returning a VoteResponse if successful.
func (s *CryptoServiceServer) Downvote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteResponse, error) {
	log.Printf("Received Downvote request for %v", in.GetSymbol())
	return &pb.VoteResponse{}, nil
}

// CreateCrypto takes a CreateReq and adds a Crypto to the database, returning
// CreateRes if successful.
func (s *CryptoServiceServer) CreateCrypto(ctx context.Context, in *pb.CreateReq) (*pb.CreateRes, error) {
	log.Printf("Received Create request for %v", in.Crypto.GetSymbol())
	return &pb.CreateRes{}, nil
}

// ReadCrypto takes a ReadReq and reads a Crypto from the database, returning
// ReadRes if successful.
func (s *CryptoServiceServer) ReadCrypto(ctx context.Context, in *pb.ReadReq) (*pb.ReadRes, error) {
	log.Printf("Received Read request for %v", in.GetQuery())
	return &pb.ReadRes{}, nil
}

// UpdateCrypto takes a UpdateReq and updates a Crypto on the database,
// returning UpdateRes if successful.
func (s *CryptoServiceServer) UpdateCrypto(ctx context.Context, in *pb.UpdateReq) (*pb.UpdateRes, error) {
	log.Printf("Received Update request for %v", in.Crypto.GetSymbol())
	return &pb.UpdateRes{}, nil
}

// DeleteCrypto takes a DeleteReq and deletes a Crypto on the database,
// returning DeleteRes if successful.
func (s *CryptoServiceServer) DeleteCrypto(ctx context.Context, in *pb.DeleteReq) (*pb.DeleteRes, error) {
	log.Printf("Received Delete request for %v", in.GetQuery())
	return &pb.DeleteRes{}, nil
}

// ListCrypto takes an Empty request, returning a stream of Crypto
func (s *CryptoServiceServer) ListCrypto(in *pb.Empty, stream pb.CryptoService_ListCryptoServer) error {
	log.Print("Received List request")
	// Test array to stream Crypto
	var testCryptos []*pb.Crypto
	testCryptos = append(testCryptos, &pb.Crypto{Id: 1, Name: "Bitcoin", Symbol: "BTC"})
	testCryptos = append(testCryptos, &pb.Crypto{Id: 2, Name: "Tether", Symbol: "USDT"})
	for _, crypto := range testCryptos {
		if err := stream.Send(crypto); err != nil {
			return err
		}
	}
	return nil
}
