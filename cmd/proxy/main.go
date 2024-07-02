package main

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
	"task1/internal/config"
)

type Application struct {
	config *config.Config
	logger *slog.Logger
}

func main() {
	cfg := *config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	logger.Info(cfg.HttpServer.Port)

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	http.ListenAndServe("localhost:8080", router)

}
