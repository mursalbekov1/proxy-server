package main

import (
	"log/slog"
	"net/http"
	"os"
	"task1/internal/config"
	"task1/internal/router"
)

func main() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	logger.Info(cfg.HttpServer.Port)
	logger.Info(cfg.HttpServer.Host)

	r := router.Router()

	err := http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		logger.Error("Error starting server: %v", err)
	}

}
