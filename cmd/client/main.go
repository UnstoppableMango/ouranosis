package main

import (
	"context"

	"github.com/charmbracelet/log"
	ouranosisv1alpha1 "github.com/unstoppablemango/ouranosis/gen/dev/unmango/ouranosis/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
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
}
