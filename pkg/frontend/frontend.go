package frontend

import (
	"io"
	"net/http"

	"github.com/unmango/go/codec"
	"github.com/unstoppablemango/ouranosis/pkg/frontend/player"
)

type Players interface {
	Post(*player.PostRequest) (*player.PostResponse, error)
}

type Frontend struct {
	Players
}

func New() *Frontend {
	return &Frontend{}
}

type HandlerFunc[T, V any] func(T) (V, error)

type RequestReader[T any] interface {
	Request(*http.Request) (T, error)
}

type ResponseWriter[T any] interface {
	Response(http.ResponseWriter, T) error
}

type Codec interface {
	Encoder(io.Writer) codec.Encoder
	Decoder(io.Reader) codec.Decoder
}

func Handler[T, V any](h HandlerFunc[T, V], c Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		dec := c.Decoder(r.Body)

		var req T
		if err := dec.Decode(&req); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		res, err := h(req)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		enc := c.Encoder(w)
		if err := enc.Encode(res); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
}
