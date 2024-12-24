package memes

import (
	"encoding/json"
	"net/http"

	"github.com/brandonbraner/memesApi/internal/http/validators"
)

type MemeQueryParams struct {
	Lat   float32 `json:"lat" schema:"lat" validate:"latitude"`
	Lon   float32 `json:"description" schema:"lon" validate:"longitude"`
	Query string  `json:"query" schema:"query"`
}

type Meme struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

func GetMeme(w http.ResponseWriter, r *http.Request) {
	var queryParams MemeQueryParams

	ok := validators.CheckQueryParamsValid(w, r, &queryParams)
	if !ok {
		return
	}

	meme := Meme{
		ID:          "1",
		URL:         "https://i.imgflip.com/30b1gx.jpg",
		Description: "Two buttons",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(meme)
	if err != nil {
		return
	}
}
