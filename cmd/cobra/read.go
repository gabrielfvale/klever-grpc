package cmd

import (
	"context"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

// readCmd represents the command to read a certain cryptocurrency
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a cryptocurrency",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create request
		req := &pb.ReadReq{Symbol: args[0]}

		// Send read request
		res, err := (*grpcClient).ReadCrypto(context.TODO(), req)
		if err != nil {
			log.Print(err)
		} else {
			log.Print(res)
		}

		return nil
	},
}
