package memes

import (
	"encoding/json"
	"log"
	"net/http"
)

type Meme struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetMeme(w http.ResponseWriter, r *http.Request) {

	meme := Meme{
		Name:        "python",
		Description: "the best language",
	}

	log.Println("MEME", meme)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(meme)
	if err != nil {
		return
	}
}
