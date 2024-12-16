package main

import (
	"github.com/chetverg999/shortener.git/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func run() error {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.PostURL)
	r.HandleFunc("/{url}", handlers.GetURL)

	return http.ListenAndServe(`:8080`, r) // запуск сервера

}
