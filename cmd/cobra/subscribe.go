package cmd

import (
	"context"
	"io"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(subscribeCmd)
}

// subscribeCmd represents the command to read a certain cryptocurrency
var subscribeCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "Read a cryptocurrency",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create request
		req := &pb.ReadReq{Symbol: args[0]}

		// Send list request, returning a stream
		stream, err := (*grpcClient).Subscribe(context.TODO(), req)

		if err != nil {
			log.Print(err)
		} else {
			// Iterate over stream
			for {
				// Get next item
				res, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("ListCrypto(_) = _, %v", err)
				}
				log.Println(res)
			}
		}

		return nil
	},
}
