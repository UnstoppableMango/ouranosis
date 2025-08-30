package player

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	ouranosis "github.com/unstoppablemango/ouranosis/pkg"
)

type Handler struct{}

type PostRequest struct {
	Name string `json:"name"`
}

// Bind implements render.Binder.
func (p *PostRequest) Bind(r *http.Request) error {
	return nil
}

type PostResponse struct {
	Player ouranosis.Player
}

// Render implements render.Renderer.
func (p *PostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handler) Post(req *PostRequest) (*PostResponse, error) {
	p := ouranosis.Player{
		Id:   uuid.New(),
		Name: req.Name,
	}

	return &PostResponse{p}, nil
}

func Create(w http.ResponseWriter, r *http.Request) {
	req := &PostRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusInternalServerError)
	} else {
		p := ouranosis.Player{
			Id:   uuid.New(),
			Name: req.Name,
		}

		render.Status(r, http.StatusCreated)
		render.Render(w, r, &PostResponse{p})
	}

	player := Player{
		Id:   uuid.New(),
		Name: p.Name,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(player); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func (h *Handler) Handle(w http.ResponseWriter, req *PostPlayerRequest) {
	p := Player{
		Id:   uuid.New(),
		Name: req.Name,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(p); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}
