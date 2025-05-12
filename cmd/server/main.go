package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	osv1alpha1 "github.com/unstoppablemango/ouranosis/gen/dev/unmango/ouranosis/v1alpha1"
	"github.com/unstoppablemango/ouranosis/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	osv1alpha1.RegisterServerServiceServer(srv, server.Server{})

	mux := http.NewServeMux()
	mux.Handle("/", srv)

	log.Info("ListenAndServing 0.0.0.0:6969")
	_ = http.ListenAndServe("0.0.0.0:6969", mux)
}
