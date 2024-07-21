package application

import (
	"encoding/json"
	"flag"
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
	data := map[string]string{
		"status":      "Available",
		"environment": app.Config.Env,
		"version":     Version,
	}
	js, err := json.Marshal(data)
	if err != nil {
		app.Logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')

	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

func (app *Application) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies", handlers.CreateMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", handlers.GetByIdHandler)
	return router
}
