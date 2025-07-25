package main

import (
	"context"
	"net"
	"os"
	"os/signal"

	"github.com/charmbracelet/log"
	osv1alpha1 "github.com/unstoppablemango/ouranosis/gen/dev/unmango/ouranosis/v1alpha1"
	"github.com/unstoppablemango/ouranosis/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:6969")
	if err != nil {
		log.Error("Failed to listen on address", "err", err)
	}

	srv := grpc.NewServer()
	osv1alpha1.RegisterServerServiceServer(srv, server.Server{})

	var cancel context.CancelFunc
	ctx := context.Background()
	ctx, cancel = signal.NotifyContext(ctx, os.Interrupt)

	go func() {
		<-ctx.Done()
		log.Info("Context completed, performing a graceful stop")
		srv.GracefulStop()
	}()

	log.Info("ListenAndServing 0.0.0.0:6969")
	if err := srv.Serve(lis); err != nil {
		log.Error("Failed serving", "err", err)
	}

	cancel()
}
