package server

import (
	"context"

	osv1alpha1 "github.com/unstoppablemango/ouranosis/gen/dev/unmango/ouranosis/v1alpha1"
)

type Server struct {
	osv1alpha1.UnimplementedServerServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Reduce(ctx context.Context, req *osv1alpha1.ReduceRequest) (*osv1alpha1.ReduceResponse, error) {
	return &osv1alpha1.ReduceResponse{
		State:  req.State,
		Events: []*osv1alpha1.ServerEvent{},
	}, nil
}
