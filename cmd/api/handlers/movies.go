package handlers

import (
	"fmt"
	"net/http"
	"watch-me-api/cmd/api/helpers"
)

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "The selected movie was: %d\n", id)
}
