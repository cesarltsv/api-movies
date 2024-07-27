package data

import (
	"time"
	"watch-me-api/cmd/api/helpers"
)

type Movie struct {
	ID        int64           `json:"id"`
	CreatedAt time.Time       `json:"-"`
	Title     string          `json:"title"`
	Year      int32           `json:"year,omitempty"`
	Runtime   helpers.Runtime `json:"runtime,omitempty"`
	Genres    []string        `json:"genres,omitempty"`
	Version   int32           `json:"version"`
}
