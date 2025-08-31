package world

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	ouranosis "github.com/unstoppablemango/ouranosis/pkg"
)

var worlds = map[uuid.UUID]*ouranosis.World{}

type PostRequest struct {
	PlayerId uuid.UUID `json:"playerId"`
}

// Bind implements render.Binder.
func (c *PostRequest) Bind(r *http.Request) error {
	return nil
}

type PostResponse struct{}

// Render implements render.Renderer.
func (c *PostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Create(w http.ResponseWriter, r *http.Request) {
	req := &PostRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	world := &ouranosis.World{}
	worlds[req.PlayerId] = world

	log.Info("Created", "world", world)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, &PostResponse{})
}
