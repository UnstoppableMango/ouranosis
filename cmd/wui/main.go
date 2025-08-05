package main

import (
	"embed"
	"io/fs"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/olivere/vite"
	"github.com/spf13/pflag"
	"github.com/unmango/go/cli"
	"github.com/unstoppablemango/ouranosis/pkg/frontend"
)

var (
	dev bool

	//go:embed all:dist
	dist embed.FS
)

func main() {
	var err error

	pflag.BoolVar(&dev, "dev", false, "Development mode")
	pflag.Parse()

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	config := vite.Config{
		IsDev:        dev,
		ViteTemplate: vite.ReactSwcTs,
	}

	if dev {
		log.Info("Running in development mode")
		config.FS = os.DirFS(".")
		config.ViteURL = "http://localhost:5173"
	} else {
		if config.FS, err = fs.Sub(dist, "dist"); err != nil {
			cli.Fail(err)
		}
	}

	viteHandler, err := vite.NewHandler(config)
	if err != nil {
		cli.Fail(err)
	}

	frontend := frontend.New()

	mux.Use(IndexHtml)
	mux.Route("/players", frontend.Players)
	mux.Get("/*", viteHandler.ServeHTTP)

	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		cli.Fail(err)
	}

	log.Info("Serving", "addr", lis.Addr())
	if err = http.Serve(lis, mux); err != nil {
		cli.Fail(err)
	}
}

func IndexHtml(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			ctx := vite.MetadataToContext(r.Context(), vite.Metadata{
				Title: "Hello, Vite!",
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// Logger will maybe one day replace middleware.Logger, we shall see
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := log.FromContext(r.Context())
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		url, t := r.Host, time.Now()

		defer func() {
			log.Info(url,
				"status", ww.Status(),
				"bytesWritten", ww.BytesWritten(),
				"header", ww.Header(),
				"elapsed", time.Since(t),
			)
		}()

		next.ServeHTTP(w, r)
	})
}
