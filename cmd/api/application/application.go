package application

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"watch-me-api/cmd/api/handlers"

	"github.com/julienschmidt/httprouter"
)

const Version = "1.0.0"

type Config struct {
	Port int
	Env  string
}

type Application struct {
	Config Config
	Logger *log.Logger
}

func New() *Application {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "Api server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (developt|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &Application{
		Config: cfg,
		Logger: logger,
	}
}

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status available")
	fmt.Fprintf(w, "environment %s\n", app.Config.Env)
	fmt.Fprintf(w, "version %s\n", Version)
}

func (app *Application) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies", handlers.CreateMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", handlers.GetByIdHandler)
	return router
}
