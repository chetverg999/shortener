package server

import (
	"fmt"
	"github.com/chetverg999/shortener.git/internal/adapter/env"
	"github.com/chetverg999/shortener.git/internal/adapter/http/handlers"
	"github.com/chetverg999/shortener.git/internal/infrastructure/mongoDB"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(registry *env.Registry, db *mongoDB.MongoDB) error {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostURL(w, r, db.Dao, registry)
	})
	r.HandleFunc("/{url}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetURL(w, r, db.Dao)
	})
	fmt.Println("Listening on port " + registry.Get("PORT") + "...")

	return http.ListenAndServe(registry.Get("PORT"), r)
}
