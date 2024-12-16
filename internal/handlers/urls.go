package handlers

import (
	"fmt"
	"github.com/chetverg999/shortener.git/internal/shortener"
	"github.com/chetverg999/shortener.git/internal/storage"
	"io"
	"net/http"
)

var BD = make(storage.Storage)

func GetURL(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[1:]
	originalURL, ok := BD[id]

	if !ok {
		http.NotFound(w, r)

		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func PostURL(w http.ResponseWriter, r *http.Request) {

	userURL, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println("Полученный url:", string(userURL))

	id := shortener.Shortener(3) // установка длины строки для сокращенной ссылки
	BD[id] = string(userURL)
	newUserURL := "http://localhost:8080/" + id

	fmt.Println("Новый url:", newUserURL)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(newUserURL))

}
