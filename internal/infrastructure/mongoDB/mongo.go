package mongoDB

import (
	"context"
	"github.com/chetverg999/shortener.git/internal/service/database"
	"github.com/chetverg999/shortener.git/internal/service/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	Client *mongo.Client
	Dao    *database.UrlDao
}

func NewMongoDB(registry *env.Registry) (*MongoDB, error) {
	opts := options.Client().ApplyURI(registry.Get("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatal(err)
	}
	dao, err := database.NewUrlDAO(registry, client)

	if err != nil {
		log.Print(err)
	}

	return &MongoDB{
		Client: client,
		Dao:    dao,
	}, nil
}
