package player

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	ouranosis "github.com/unstoppablemango/ouranosis/pkg"
)

type Handler struct {
	f ouranosis.Frontend
}

func NewHandler(f ouranosis.Frontend) *Handler {
	return &Handler{f}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)

	p := &PostPlayerRequest{}
	if err := dec.Decode(p); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
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
