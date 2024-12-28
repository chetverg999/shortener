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

func startApp() error {
	param := env.GetRegistry()
	mongoDB, err := storage.NewMongoDB(param)

	if err != nil {
		log.Print(err)
	}
	defer func() {

		if err := mongoDB.Client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Ошибка при отключении от MongoDB: %v", err)
		}
		fmt.Println("Соединение с MongoDB закрыто.")
	}()
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostURL(w, r, mongoDB.Dao)
	})
	r.HandleFunc("/{url}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetURL(w, r, mongoDB.Dao)
	})
	fmt.Println("Listening on port " + param.Get("PORT") + "...")

	return http.ListenAndServe(param.Get("PORT"), r)
}
