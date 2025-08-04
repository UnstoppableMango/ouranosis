package player

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type PostPlayerRequest struct {
	Name string `json:"name"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()

	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)

	p := &PostPlayerRequest{}
	if err := dec.Decode(p); err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		Handle(w, p)
	}
}

func Handle(w http.ResponseWriter, req *PostPlayerRequest) {
	p := Player{
		Id:   uuid.New(),
		Name: req.Name,
	}
}

type Player struct {
	Id   uuid.UUID
	Name string
}
