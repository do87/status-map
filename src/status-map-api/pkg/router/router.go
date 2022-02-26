package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// New returns a new chi router
func New(isLocal bool) *chi.Mux {
	r := chi.NewRouter()

	r.Use(getCorsHandler(isLocal))
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.RedirectSlashes)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	return r
}

func getCorsHandler(isLocal bool) func(next http.Handler) http.Handler {
	opts := cors.Options{
		AllowedOrigins:   []string{"https://odj-status-map.prod.odjui.sys.odj.cloud"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}

	if isLocal {
		opts.AllowedOrigins = []string{"http://*"}
	}
	return cors.Handler(opts)
}
