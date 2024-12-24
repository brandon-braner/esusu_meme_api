package v1

import (
	"github.com/brandonbraner/memesApi/internal/v1/memes"
	"net/http"
)

func Setup() *http.ServeMux {
	v1Router := http.NewServeMux()

	v1Router.HandleFunc("GET /health_check", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Health Check v1"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	memesRouter := memes.Setup()
	v1Router.Handle("/memes/", http.StripPrefix("/memes", memesRouter))

	return v1Router
}
