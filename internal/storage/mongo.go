package storage

import (
	"context"
	"github.com/chetverg999/shortener.git/internal/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	Client *mongo.Client
	Dao    *UrlDao
}

func NewMongoDB(registry *env.Registry) (*MongoDB, error) {
	ctx := context.Background()
	opts := options.Client().ApplyURI(registry.Get("DB_HOST"))
	client, err := mongo.Connect(ctx, opts)

	if err != nil {

		panic(err)
	}
	dao, err := NewUrlDAO(registry, client)

	if err != nil {
		log.Print(err)
	}

	return &MongoDB{
		Client: client,
		Dao:    dao,
	}, nil
}
