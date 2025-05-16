package client

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	ouranosisv1alpha1 "github.com/unstoppablemango/ouranosis/gen/dev/unmango/ouranosis/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTextCommand() *cobra.Command {
	return &cobra.Command{
		Use: "text",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("0.0.0.0:6969",
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				log.Error("Failed to create server connection", "err", err)
			}

			client := ouranosisv1alpha1.NewServerServiceClient(conn)

			ctx := context.Background()
			res, err := client.Reduce(ctx, &ouranosisv1alpha1.ReduceRequest{})
			if err != nil {
				log.Error("Failed to reduce from server", "err", err)
			} else {
				log.Info("Got response", "res", res)
			}
		},
	}
}
