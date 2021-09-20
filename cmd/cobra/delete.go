package cmd

import (
	"context"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

// deleteCmd represents the command to delete a certain cryptocurrency
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a cryptocurrency",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create request
		req := &pb.DeleteReq{Symbol: args[0]}

		// Send create request
		res, err := (*grpcClient).DeleteCrypto(context.TODO(), req)
		if err != nil {
			log.Print(err)
		} else {
			log.Print(res)
		}

		return nil
	},
}
