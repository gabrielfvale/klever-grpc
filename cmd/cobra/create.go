package cmd

import (
	"context"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

func init() {
	createCmd.Flags().StringP("symbol", "s", "", "The cryptocurrency symbol")
	createCmd.Flags().StringP("name", "n", "", "The cryptocurrency name")
	createCmd.MarkFlagRequired("symbol")
	createCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(createCmd)
}

// createCmd represents the command to create a new cryptocurrency
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Add a new cryptocurrency",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Get symbol from flag
		symbol, err := cmd.Flags().GetString("symbol")
		if err != nil {
			return err
		}
		// Get name from flag
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		// Create request
		req := &pb.CreateReq{Crypto: &pb.Crypto{
			Name:   name,
			Symbol: symbol,
		}}

		// Send create request
		res, err := (*grpcClient).CreateCrypto(context.TODO(), req)
		if err != nil {
			log.Print(err)
		} else {
			log.Print(res)
		}

		return nil
	},
}
