package memes

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/schema"
	"net/http"
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
	// start request validation
	// TODO: There has to be a better way to do this
	validate := validator.New()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request 1", http.StatusBadRequest)
		return
	}
	var queryParams MemeQueryParams
	decoder := schema.NewDecoder()

	decoder.IgnoreUnknownKeys(true)
	err = decoder.Decode(&queryParams, r.Form)

	if err != nil {
		http.Error(w, "Bad Request 2", http.StatusBadRequest)
		return
	}

	err = validate.Struct(queryParams)
	if err != nil {
		// Handle validation errors
		if _, ok := err.(*validator.InvalidValidationError); ok {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		errorMessages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' validation failed on tag '%s'", err.Field(), err.Tag()))
		}
		http.Error(w, fmt.Sprintf("Validation failed: %s", errorMessages), http.StatusBadRequest)
		return
	}
	// end request validation
	meme := Meme{
		ID:          "1",
		URL:         "https://i.imgflip.com/30b1gx.jpg",
		Description: "Two buttons",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(meme)
	if err != nil {
		return
	}
}
