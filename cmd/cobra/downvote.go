package cmd

import (
	"context"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downvoteCmd)
}

// downvoteCmd represents the command to downvote a certain cryptocurrency
var downvoteCmd = &cobra.Command{
	Use:   "downvote",
	Short: "Downvote a cryptocurrency",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create request
		req := &pb.VoteRequest{Symbol: args[0]}

		// Send downvote request
		res, err := (*grpcClient).Downvote(context.TODO(), req)
		if err != nil {
			log.Print(err)
		} else {
			log.Print(res)
		}

		return nil
	},
}
