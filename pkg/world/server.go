package world

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	osv1alpha1 "github.com/unstoppablemango/ouranosis/gen/dev/unmango/ouranosis/v1alpha1"
	ouranosis "github.com/unstoppablemango/ouranosis/pkg"
)

type Server struct {
	osv1alpha1.UnimplementedWorldServiceServer

	clients []string

	World *ouranosis.World
}

// Join implements ouranosisv1alpha1.WorldServiceServer.
func (s *Server) Join(ctx context.Context, req *osv1alpha1.JoinRequest) (*osv1alpha1.JoinResponse, error) {
	log := log.FromContext(ctx)
	uuid := uuid.New().String()

	log.Info("Client joining", "uuid", uuid)
	s.clients = append(s.clients, uuid)

	return &osv1alpha1.JoinResponse{
		Uuid: uuid,
	}, nil
}
