package handlers

import (
	"fmt"
	"github.com/chetverg999/shortener.git/internal/env"
	"github.com/chetverg999/shortener.git/internal/shortener"
	"github.com/chetverg999/shortener.git/internal/storage"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

const shortUrlLength = 3

func GetURL(w http.ResponseWriter, r *http.Request, collection *storage.UrlDao) {
	short := r.URL.Path[1:]
	originalURL, err := collection.Find(short)

	if err != nil {
		http.NotFound(w, r)
		fmt.Println(err)

		return
	}
	http.Redirect(w, r, originalURL.UserURL, http.StatusFound)
}

func PostURL(w http.ResponseWriter, r *http.Request, collection *storage.UrlDao) {
	userURL, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	err = r.Body.Close()

	if err != nil {

		return
	}

	if parseURL(userURL) != nil { // Валидирование
		http.Error(w, parseURL(userURL).Error(), http.StatusBadRequest)

		return
	}
	fmt.Println("Полученный url:", string(userURL))
	shortURL := shortener.Shortener(shortUrlLength)
	data := &storage.ShortURL{
		Id:      bson.NewObjectId(),
		UserURL: string(userURL),
		Short:   shortURL,
	}
	err = collection.Insert(data)

	if err != nil {
		fmt.Println(err)

		return
	}
	newUserURL := env.GoDotEnvVariable("HOST") + shortURL
	fmt.Println("Новый url:", newUserURL)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(newUserURL))
	if err != nil {

		return
	}
}
