package cmd

import (
	"fmt"
	"os"

	pb "github.com/gabrielfvale/klever-grpc/internal/proto"
	"github.com/spf13/cobra"
)

var grpcClient *pb.CryptoServiceClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Upvote and downvote well-known cryptocurrencies",
}

// Execute runs the rootCmd, checking for errors
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// SetClient sets the global grpcClient variable, used to access client methods
func SetClient(client *pb.CryptoServiceClient) {
	grpcClient = client
}
