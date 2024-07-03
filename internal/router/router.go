package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"task1/internal/proxy"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthCheck", proxy.HealthCheck)
		r.Get("/external-api", proxy.Proxy)
	})

	return router
}
