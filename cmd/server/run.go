package main

import (
	"github.com/chetverg999/shortener.git/internal/handlers"
	"net/http"
)

func run() error {

	mux := http.NewServeMux()
	//mux.HandleFunc("/", handlers.PostURL)
	//mux.HandleFunc("/{id}", handlers.GetURL)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handlers.PostURL(w, r)
		} else {
			handlers.GetURL(w, r)
		}
	})

	return http.ListenAndServe(`:8080`, mux) // запуск сервера

}
