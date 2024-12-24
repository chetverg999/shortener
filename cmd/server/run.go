package main

import (
	"context"
	"fmt"
	"github.com/chetverg999/shortener.git/internal/env"
	"github.com/chetverg999/shortener.git/internal/handlers"
	"github.com/chetverg999/shortener.git/internal/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func run() error {

	client, urlCollection := storage.StartMongo() // подключение к базе данных
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Ошибка при отключении от MongoDB: %v", err)
		}
		fmt.Println("Соединение с MongoDB закрыто.")
	}()
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostURL(w, r, urlCollection)
	})
	r.HandleFunc("/{url}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetURL(w, r, urlCollection)
	})
	port := env.GoDotEnvVariable("PORT")
	fmt.Println("Listening on port " + port + "...")

	return http.ListenAndServe(port, r)
}
