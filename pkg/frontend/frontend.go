package frontend

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Frontend struct {
	State *State
}

func New() *Frontend {
	return &Frontend{}
}

func (f Frontend) Players(r chi.Router) {
	r.Post("/", f.CreatePlayer)
}

func (f Frontend) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	req := &CreatePlayerRequest{}
	if err := dec.Decode(req); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	res, err := f.State.PostPlayer(r.Context(), req)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	enc := json.NewEncoder(w)
	if err = enc.Encode(res); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
