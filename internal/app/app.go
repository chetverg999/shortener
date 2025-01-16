package app

import (
	"context"
	"github.com/chetverg999/shortener.git/internal/infrastructure/mongoDB"
	"github.com/chetverg999/shortener.git/internal/infrastructure/server"
	"github.com/chetverg999/shortener.git/internal/service/env"
	"log"
)

func Run() error {
	registry := env.NewRegistry()
	db, err := mongoDB.NewMongoDB(registry)

	if err != nil {
		log.Print(err)
	}
	defer func() {

		if err := db.Client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Ошибка при отключении от MongoDB: %v", err)
		}
	}()

	return server.NewRouter(registry, db)
}
