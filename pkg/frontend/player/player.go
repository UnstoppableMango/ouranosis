package player

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	ouranosis "github.com/unstoppablemango/ouranosis/pkg"
)

var players = map[uuid.UUID]ouranosis.Player{}

type Handler struct{}

type GetResponse struct {
	Player ouranosis.Player `json:"player"`
}

// Render implements render.Renderer.
func (g *GetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type PostRequest struct {
	Name string `json:"name"`
}

// Bind implements render.Binder.
func (p *PostRequest) Bind(r *http.Request) error {
	return nil
}

type PostResponse struct {
	Player ouranosis.Player `json:"player"`
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

		players[p.Id] = p

		log.Info("Created", "player", p)
		render.Status(r, http.StatusCreated)
		render.Render(w, r, &PostResponse{p})
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		return
	}

	p, ok := players[id]
	if !ok {
		log.Info("Not found", "id", id)
		render.Status(r, http.StatusNotFound)
		return
	}

	log.Info("Found", "player", p)
	render.Status(r, http.StatusOK)
	render.Render(w, r, &GetResponse{p})
}
