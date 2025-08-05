package frontend

import "github.com/go-chi/chi/v5"

type Rest interface {
	Players(chi.Router)
}
