package impl

import (
	"context"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/gabrielfvale/klever-grpc/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CryptoType represents the bson readable data from the protobuf
type CryptoType struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Symbol    string             `bson:"symbol"`
	Name      string             `bson:"name"`
	Upvotes   int32              `bson:"upvotes"`
	Downvotes int32              `bson:"downvotes"`
}

// ChangeEvent represents a fraction of the change stream event returned by
// collection.Watch()
type ChangeEvent struct {
	FullDocument CryptoType `bson:"fullDocument"`
}

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
	readCrypto := CryptoType{}

	// Load collection
	collection, err := pkg.GetMongoCollection()
	if err != nil {
		return nil, err
	}

	// Read crypto of a given symbol
	res := collection.FindOne(context.TODO(), bson.M{"symbol": in.GetSymbol()})
	if err := res.Decode(&readCrypto); err != nil {
		return nil, status.Error(codes.NotFound, "Could not find Object")
	}

	// Update object
	readCrypto.Upvotes += 1
	_, err = collection.ReplaceOne(context.TODO(), primitive.M{"_id": readCrypto.Id}, readCrypto)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update crypto: %v", err)
	}

	return &pb.VoteResponse{
		Name:   readCrypto.Name,
		Symbol: readCrypto.Symbol,
		Votes:  readCrypto.Upvotes - readCrypto.Downvotes,
	}, nil
}

// Downvote takes a VoteRequest and updates the "downvotes" field on a given crypto
// returning a VoteResponse if successful.
func (s *CryptoServiceServer) Downvote(ctx context.Context, in *pb.VoteRequest) (*pb.VoteResponse, error) {
	log.Printf("Received Downvote request for %v", in.GetSymbol())
	readCrypto := CryptoType{}

	// Load collection
	collection, err := pkg.GetMongoCollection()
	if err != nil {
		return nil, err
	}

	// Read crypto of a given symbol
	res := collection.FindOne(context.TODO(), bson.M{"symbol": in.GetSymbol()})
	if err := res.Decode(&readCrypto); err != nil {
		return nil, status.Error(codes.NotFound, "Could not find Object")
	}

	// Update object
	readCrypto.Downvotes += 1
	_, err = collection.ReplaceOne(context.TODO(), primitive.M{"_id": readCrypto.Id}, readCrypto)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update crypto: %v", err)
	}

	return &pb.VoteResponse{
		Name:   readCrypto.Name,
		Symbol: readCrypto.Symbol,
		Votes:  readCrypto.Upvotes - readCrypto.Downvotes,
	}, nil
}

// CreateCrypto takes a CreateReq and adds a Crypto to the database, returning
// CreateRes if successful.
func (s *CryptoServiceServer) CreateCrypto(ctx context.Context, in *pb.CreateReq) (*pb.CreateRes, error) {
	log.Printf("Received Create request for %v", in.Crypto.GetSymbol())
	crypto := in.GetCrypto()

	// Load collection
	collection, err := pkg.GetMongoCollection()
	if err != nil {
		return nil, err
	}

	// Check if crypto of "symbol" aready exists
	query := collection.FindOne(context.TODO(), bson.M{"symbol": crypto.GetSymbol()})
	if err := query.Decode(&CryptoType{}); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Crypto of symbol %s already exists", crypto.GetSymbol())
	}

	item := CryptoType{
		Symbol:    crypto.GetSymbol(),
		Name:      crypto.GetName(),
		Upvotes:   crypto.GetUpvotes(),
		Downvotes: crypto.GetDownvotes(),
	}

	// Insert crypto item
	res, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error on InsertOne: %v", err)
	}
	// Print inserted id
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Could not convert Id to oid")
	}
	log.Printf("Inserted Crypto of ID %s", oid.Hex())

	return &pb.CreateRes{Crypto: crypto}, nil
}

// ReadCrypto takes a ReadReq and reads a Crypto from the database, returning
// ReadRes if successful.
func (s *CryptoServiceServer) ReadCrypto(ctx context.Context, in *pb.ReadReq) (*pb.ReadRes, error) {
	log.Printf("Received Read request for %v", in.GetSymbol())
	item := CryptoType{}

	// Load collection
	collection, err := pkg.GetMongoCollection()
	if err != nil {
		return nil, err
	}

	// Read crypto of a given symbol
	res := collection.FindOne(context.TODO(), bson.M{"symbol": in.GetSymbol()})
	if err := res.Decode(&item); err != nil {
		return nil, status.Error(codes.NotFound, "Could not find Object")
	}

	log.Printf("Read crypto %s", item.Name)

	response := &pb.ReadRes{
		Crypto: &pb.Crypto{
			Id:        item.Id.Hex(),
			Symbol:    item.Symbol,
			Name:      item.Name,
			Upvotes:   item.Upvotes,
			Downvotes: item.Downvotes,
		},
	}

	return response, nil
}

// UpdateCrypto takes a UpdateReq and updates a Crypto on the database,
// returning UpdateRes if successful.
func (s *CryptoServiceServer) UpdateCrypto(ctx context.Context, in *pb.UpdateReq) (*pb.UpdateRes, error) {
	log.Printf("Received Update request for %v", in.Crypto.GetSymbol())
	crypto := in.GetCrypto()
	updtCrypto := CryptoType{}

	// Load collection
	collection, err := pkg.GetMongoCollection()
	if err != nil {
		return nil, err
	}

	// Read crypto of a given symbol
	res := collection.FindOne(context.TODO(), bson.M{"symbol": crypto.GetSymbol()})
	if err := res.Decode(&updtCrypto); err != nil {
		return nil, status.Error(codes.NotFound, "Could not find Object")
	}

	// Only update if necessary
	if crypto.Name != "" {
		updtCrypto.Name = crypto.Name
	}
	if crypto.Upvotes != -1 {
		updtCrypto.Upvotes = crypto.Upvotes
	}
	if crypto.Downvotes != -1 {
		updtCrypto.Downvotes = crypto.Downvotes
	}

	// Update object
	_, err = collection.ReplaceOne(context.TODO(), primitive.M{"_id": updtCrypto.Id}, updtCrypto)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update crypto: %v", err)
	}

	return &pb.UpdateRes{Crypto: &pb.Crypto{
		Symbol:    crypto.Symbol,
		Name:      updtCrypto.Name,
		Upvotes:   updtCrypto.Upvotes,
		Downvotes: updtCrypto.Downvotes,
	}}, nil
}

// DeleteCrypto takes a DeleteReq and deletes a Crypto on the database,
// returning DeleteRes if successful.
func (s *CryptoServiceServer) DeleteCrypto(ctx context.Context, in *pb.DeleteReq) (*pb.DeleteRes, error) {
	log.Printf("Received Delete request for %v", in.GetSymbol())

	// Load collection
	collection, err := pkg.GetMongoCollection()
	if err != nil {
		return nil, err
	}
	// Try to delete crypto of given symbol
	res, err := collection.DeleteOne(context.TODO(), bson.M{"symbol": in.GetSymbol()})
	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		return &pb.DeleteRes{Ok: false}, status.Errorf(codes.Internal, "Cannot delete Crypto: %v")
	}

	return &pb.DeleteRes{Ok: true}, nil
}

// ListCrypto takes an Empty request, returning a stream of Crypto
func (s *CryptoServiceServer) ListCrypto(in *pb.Empty, stream pb.CryptoService_ListCryptoServer) error {
	log.Print("Received List request")
	// Load collection
	collection, err := pkg.GetMongoCollection()
	if err != nil {
		return err
	}
	// Create Mongo cursor
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return status.Errorf(codes.Internal, "Error: %v", err)
	}
	defer cursor.Close(context.TODO())
	// Iterate over cursor entries
	for cursor.Next(context.TODO()) {
		item := &CryptoType{}
		if err := cursor.Decode(item); err != nil {
			return status.Errorf(codes.Internal, "Could not decode data: %v", err)
		}
		stream.Send(&pb.Crypto{
			Symbol:    item.Symbol,
			Name:      item.Name,
			Upvotes:   item.Upvotes,
			Downvotes: item.Downvotes,
		})
	}
	// Check for cursor errors
	if err = cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "Error: %v", err)
	}
	return nil
}

// Subscribe creates a real time stream that waits for changes in the database
// with a given symbol filter.
func (s *CryptoServiceServer) Subscribe(in *pb.ReadReq, stream pb.CryptoService_SubscribeServer) error {
	log.Printf("Received Subscribe request for %v", in.GetSymbol())

	// Load collection
	collection, err := pkg.GetMongoCollection()
	if err != nil {
		return err
	}

	// Set pipeline filter
	pipeline := mongo.Pipeline{
		bson.D{{
			Key:   "$match",
			Value: bson.D{{Key: "fullDocument.symbol", Value: in.GetSymbol()}}},
		},
	}
	streamOptions := options.ChangeStream().SetFullDocument(options.UpdateLookup)

	// Create a change stream
	mstream, err := collection.Watch(context.TODO(), pipeline, streamOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Wait for events and stream to client
	var event ChangeEvent
	for mstream.Next(context.TODO()) {
		if err := mstream.Decode(&event); err != nil {
			log.Fatalf("Error decoding: %v", err)
		}
		crypto := event.FullDocument
		stream.Send(&pb.Crypto{
			Symbol:    crypto.Symbol,
			Name:      crypto.Name,
			Upvotes:   crypto.Upvotes,
			Downvotes: crypto.Downvotes,
		})
	}
	return nil
}
