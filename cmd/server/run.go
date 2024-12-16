package main

import (
	"fmt"
	"github.com/chetverg999/shortener.git/internal/env"
	"github.com/chetverg999/shortener.git/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func run() error {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.PostURL)
	r.HandleFunc("/{url}", handlers.GetURL)
	port := env.GoDotEnvVariable("PORT")
	fmt.Println("Listening on port " + port + "...")

	return http.ListenAndServe(port, r)
}
