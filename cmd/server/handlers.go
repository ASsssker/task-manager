package server

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
)

// Ping server connect
func (app *Applicaton) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func (app *Applicaton) GetTasks(w http.ResponseWriter, r *http.Request) {
	data, err := app.TaskService.GetTasks()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write(data.Bytes())
}

func (app *Applicaton) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_conv, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	data, err := app.TaskService.Get(uint(id_conv))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
		return
	}

	w.Write(data.Bytes())
}

func (app *Applicaton) PostTask(w http.ResponseWriter, r *http.Request) {
	data, err := app.TaskService.Insert(r.Body)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.WriteHeader(http.StatusCreated)
	w.Write(data.Bytes())
}

func (app *Applicaton) PutTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_conv, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	data, err := app.TaskService.Update(uint(id_conv), r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	w.Write(data.Bytes())
}