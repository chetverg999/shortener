package app

import (
	"context"
	"github.com/chetverg999/shortener.git/internal/adapter/env"
	"github.com/chetverg999/shortener.git/internal/infrastructure/mongoDB"
	"github.com/chetverg999/shortener.git/internal/infrastructure/server"
	"log"
)

func Run() error {
	registry := env.GetRegistry()
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
