package memes

import (
	"net/http"
)

func Setup() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /health_check", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Health Check memes"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	router.HandleFunc("GET /", GetMeme)

	return router
}
