package cmd

import (
	"context"
	"log"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

func init() {
	updateCmd.Flags().StringP("symbol", "s", "", "The cryptocurrency symbol")
	updateCmd.Flags().StringP("name", "n", "", "The cryptocurrency name")
	updateCmd.Flags().Int32P("upvotes", "u", -1, "The amount of upvotes")
	updateCmd.Flags().Int32P("downvotes", "d", -1, "The amount of downvotes")
	updateCmd.MarkFlagRequired("symbol")
	rootCmd.AddCommand(updateCmd)
}

// updateCmd represents the command to update a certain cryptocurrency
var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"updt"},
	Short:   "Update a cryptocurrency",
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
		// Get upvotes from flag
		upvotes, err := cmd.Flags().GetInt32("upvotes")
		if err != nil {
			return err
		}
		// Get downvotes from flag
		downvotes, err := cmd.Flags().GetInt32("downvotes")
		if err != nil {
			return err
		}

		// Create request
		req := &pb.UpdateReq{Crypto: &pb.Crypto{
			Name:      name,
			Symbol:    symbol,
			Upvotes:   upvotes,
			Downvotes: downvotes,
		}}

		// Send create request
		res, err := (*grpcClient).UpdateCrypto(context.TODO(), req)
		if err != nil {
			log.Print(err)
		} else {
			log.Print(res)
		}

		return nil
	},
}
