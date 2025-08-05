package frontend

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/unstoppablemango/ouranosis/pkg/frontend/player"
)

type Frontend struct {
	players chan player.Player
}

func New() *Frontend {
	return &Frontend{
		players: make(chan player.Player),
	}
}

func (f *Frontend) Players(r chi.Router) {
	p := player.NewHandler(f)
	r.Post("/", p.Post)
}

func (f *Frontend) Post(w http.ResponseWriter, r *http.Request) {

}
