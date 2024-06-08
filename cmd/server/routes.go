package server

import "github.com/go-chi/chi/v5"

func (app *Applicaton) geRoutes() *chi.Mux {
	r := chi.NewMux()
	r.Get("/", app.Ping)
	
	return r
}