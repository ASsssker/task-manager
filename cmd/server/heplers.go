package server

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Applicaton) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Applicaton) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Applicaton) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}