package server

import (
	"context"

	osv1alpha1 "github.com/unstoppablemango/ouranosis/gen/dev/unmango/ouranosis/v1alpha1"
)

type Server struct {
	osv1alpha1.UnimplementedServerServiceServer
}

func (Server) Reduce(context.Context, *osv1alpha1.ReduceRequest) (*osv1alpha1.ReduceResponse, error) {
	return nil, nil
}
