package customerrors

import (
	"fmt"
	"net/http"
	"watch-me-api/cmd/api/helpers"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := helpers.Envelop{"error": message}
	err := helpers.WriteJson(w, status, env, nil)
	if err != nil {
		fmt.Println("Errors: ", err)
		fmt.Println("Request: ", r)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request"
	ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	ErrorResponse(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not suppoert for this resource", r.Method)
	ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}
