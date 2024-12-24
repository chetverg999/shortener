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
	}

	if parseURL(userURL) != nil {
		http.Error(w, parseURL(userURL).Error(), http.StatusBadRequest)
	} else {
		fmt.Println("Полученный url:", string(userURL))

		short := shortener.Shortener(3) // установка длины строки для сокращенной ссылки

		data := &storage.ShortURL{
			Id:      bson.NewObjectId(),
			UserURL: string(userURL),
			Short:   short,
		}

		err = collection.Insert(data)

		if err != nil {
			fmt.Println(err)
		}

		newUserURL := env.GoDotEnvVariable("HOST") + short

		fmt.Println("Новый url:", newUserURL)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(newUserURL))
	}
}
