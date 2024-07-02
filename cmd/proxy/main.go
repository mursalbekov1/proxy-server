package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"task1/internal/config"
)

func main() {
	cfg := config.MustLoad()

	log.Println(cfg.Port)

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe("localhost:8080", router)

}
