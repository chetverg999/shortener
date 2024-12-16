package handlers

import (
	"fmt"
	"github.com/chetverg999/shortener.git/internal/shortener"
	"github.com/chetverg999/shortener.git/internal/storage"
	"html/template"
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

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка при разборе формы", http.StatusBadRequest)
		return
	}

	templ, err := template.ParseFiles("internal/template/start.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusMovedPermanently)
		return
	}

	userURL := r.Form.Get("url")

	id := shortener.Shortener()
	BD[id] = userURL
	newUserURL := "http://localhost:8080/" + id
	fmt.Println("Новый url:", newUserURL)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	data := map[string]interface{}{
		"ShortURL": newUserURL,
	}

	if userURL == "" {
		err = templ.Execute(w, nil)
		if err != nil {
			http.Error(w, "Ошибка рендеринга шаблона", http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("Полученный url:", userURL)
		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, "Ошибка рендеринга шаблона", http.StatusInternalServerError)
			return
		}

	}
}
