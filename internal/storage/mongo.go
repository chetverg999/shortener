package storage

import (
	"context"
	"fmt"
	"github.com/chetverg999/shortener.git/internal/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func StartMongo() (*mongo.Client, *UrlDao) {
	ctx := context.TODO()
	opts := options.Client().ApplyURI(env.GoDotEnvVariable("DB"))

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", client)

	ShortenUrlDAO, err := NewUrlDAO(ctx, client)
	if err != nil {
		log.Print(err)
	}

	return client, ShortenUrlDAO
}
