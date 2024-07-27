package application

import (
	"flag"
	"log"
	"net/http"
	"os"
	customerrors "watch-me-api/cmd/api/customErrors"
	"watch-me-api/cmd/api/handlers"
	"watch-me-api/cmd/api/helpers"

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
	data := helpers.Envelop{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.Config.Env,
			"version":     Version,
		},
	}

	err := helpers.WriteJson(w, http.StatusOK, data, nil)

	if err != nil {
		app.Logger.Println(err)
		customerrors.ServerErrorResponse(w, r, err)
	}
}

func (app *Application) Routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(customerrors.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(customerrors.MethodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies", handlers.CreateMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", handlers.GetByIdHandler)
	return router
}

func (app *Application) LogError(r *http.Request, err error) {
	app.Logger.Println(err)
}
