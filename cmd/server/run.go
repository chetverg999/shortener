package main

import (
	"github.com/chetverg999/shortener.git/internal/handlers"
	"net/http"
)

func run() error {
	return http.ListenAndServe(`:8080`, http.HandlerFunc(handlers.Start)) // запуск сервера
}
