package player

import (
	"encoding/json"
	"net/http"

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
	} else {
		Handle(w, p)
	}
}
