package cmd

import (
	"context"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(upvoteCmd)
}

// upvoteCmd represents the command to upvote a certain cryptocurrency
var upvoteCmd = &cobra.Command{
	Use:     "upvote",
	Aliases: []string{"up"},
	Short:   "Upvote a cryptocurrency",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create request
		req := &pb.VoteRequest{Symbol: args[0]}

		// Send upvote request
		res, err := (*grpcClient).Upvote(context.TODO(), req)
		if err != nil {
			log.Print(err)
		} else {
			log.Print(res)
		}

		return nil
	},
}
