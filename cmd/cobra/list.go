package cmd

import (
	"context"
	"io"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

// listCmd represents the command to list the added cryptocurrencies
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List the cryptocurrencies",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Send list request, returning a stream
		stream, err := (*grpcClient).ListCrypto(context.TODO(), &pb.Empty{})

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
