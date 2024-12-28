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

const (
	shortUrlLength     = 3
	mediaType          = "Content-Type"
	mediaTypeTextPlain = "text/plain; charset=utf-8"
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

func PostURL(w http.ResponseWriter, r *http.Request, collection *storage.UrlDao, registry *env.Registry) {
	userURL, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = r.Body.Close(); err != nil {

		return
	}

	if parseURL(userURL) != nil { // Валидирование
		http.Error(w, parseURL(userURL).Error(), http.StatusBadRequest)

		return
	}
	shortURL := shortener.Shortener(shortUrlLength)
	data := &storage.ShortURL{
		Id:      bson.NewObjectId(),
		UserURL: string(userURL),
		Short:   shortURL,
	}

	if err = collection.Insert(data); err != nil {
		fmt.Println(err)

		return
	}
	newUserURL := registry.Get("HOST") + shortURL
	w.Header().Set(mediaType, mediaTypeTextPlain)
	w.WriteHeader(http.StatusCreated)

	if _, err = w.Write([]byte(newUserURL)); err != nil {

		return
	}
}
