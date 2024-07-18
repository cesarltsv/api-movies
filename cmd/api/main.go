package main

import (
	"fmt"
	"net/http"
	"time"
	"watch-me-api/cmd/api/application"
)

func main() {

	app := application.New()

	logger := app.Logger
	cfg := app.Config

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
