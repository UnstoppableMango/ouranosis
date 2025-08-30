package frontend

import (
	"context"

	"github.com/google/uuid"
	ouranosis "github.com/unstoppablemango/ouranosis/pkg"
)

type State struct {
	Players []ouranosis.Player
}

type CreatePlayerRequest struct {
	Name string `json:"name"`
}

type CreatePlayerResponse struct {
	Player ouranosis.Player `json:"player"`
}

func (s *State) PostPlayer(ctx context.Context, req *CreatePlayerRequest) (*CreatePlayerResponse, error) {
	player := ouranosis.Player{
		Id:   uuid.New(),
		Name: req.Name,
	}

	s.Players = append(s.Players, player)

	return &CreatePlayerResponse{
		Player: player,
	}, nil
}

type ListPlayersRequest struct{}

type ListPlayersResponse struct {
	Players []ouranosis.Player
}

func (s *State) ListPlayers(ctx context.Context, req *ListPlayersRequest) (*ListPlayersResponse, error) {
	return &ListPlayersResponse{
		Players: s.Players,
	}, nil
}
