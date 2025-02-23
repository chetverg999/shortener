package database

import (
	"context"
	"fmt"
	"github.com/chetverg999/shortener.git/internal/entity"
	"github.com/chetverg999/shortener.git/internal/service/env"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"unicode/utf8"
)

type UrlDao struct {
	c *mongo.Collection
}

func NewUrlDAO(registry *env.Registry, client *mongo.Client) (*UrlDao, error) {
	name := registry.Get("DB_NAME")
	dbCollection := registry.Get("DB_COLLECTION")

	return &UrlDao{
		c: client.Database(name).Collection(dbCollection),
	}, nil
}

func countDocuments(collection *UrlDao) (int64, error) {
	count, err := collection.c.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *UrlDao) Insert(url *entity.ShortURL) error {
	if !utf8.ValidString(url.UserURL) {
		fmt.Println("Некорректная строка UTF-8:", url.UserURL)
	}
	beforeCount, _ := countDocuments(u)
	fmt.Println("Данные для вставки:")
	fmt.Printf("Id: %v\n", url.Id)
	fmt.Printf("UserURL: %s\n", url.UserURL)
	fmt.Printf("Short: %s\n", url.Short)
	_, err := u.c.InsertOne(context.Background(), url)

	if err != nil {
		fmt.Println("Ошибка при вставке в MongoDB:", err)
	}
	afterCount, _ := countDocuments(u)

	if afterCount > beforeCount {
		fmt.Println("Документ успешно записан в базу данных.")
	} else {
		fmt.Println("Документ не был записан.")
	}

	return err
}

func (u *UrlDao) Find(id string) (*entity.ShortURL, error) {
	filter := bson.M{"Short": id}
	fmt.Println("Идет поиск в базе")
	var shortURL entity.ShortURL
	err := u.c.FindOne(context.Background(), filter).Decode(&shortURL)
	fmt.Println(err)
	switch {
	case err == mongo.ErrNoDocuments:
		fmt.Println("Документ не был найден.")
		return nil, mux.ErrNotFound
	case err == nil:
		fmt.Println("Документ найден")
		return &shortURL, nil
	default:
		fmt.Println("Неизвестная ошибка")
		return nil, err
	}
}
