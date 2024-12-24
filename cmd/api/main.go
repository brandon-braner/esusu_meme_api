package main

import (
	"fmt"
	"github.com/brandonbraner/memesApi/internal/v1"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /health_check", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Health Check"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	// setup v1 routes
	v1Router := v1.Setup()
	router.Handle("/v1/", http.StripPrefix("/v1", v1Router))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
