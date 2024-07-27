package handlers

import (
	"fmt"
	"net/http"
	"time"
	customerrors "watch-me-api/cmd/api/customErrors"
	"watch-me-api/cmd/api/helpers"
	"watch-me-api/internals/data"
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

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Madacascar",
		Runtime:   102,
		Genres:    []string{"drama", "comedy"},
		Version:   1,
	}

	err = helpers.WriteJson(w, http.StatusOK, helpers.Envelop{"movie": movie}, nil)
	if err != nil {
		customerrors.ServerErrorResponse(w, r, err)
	}

}
