package server

import "github.com/go-chi/chi/v5"

func (app *Applicaton) getRoutes() *chi.Mux {
	r := chi.NewMux()
	
	r.Use(app.RequestsInfo, app.CompressResponse, app.SetApplicationHeader)

	r.Get("/", app.Ping)
	
	r.Route("/api/tasks", func(r chi.Router) {
		r.Get("/", app.GetTasks)
		r.Post("/", app.PostTask)
		r.Get("/{id}", app.GetTask)
		r.Put("/{id}", app.PutTask)
		r.Delete("/{id}", app.DeleteTask)
	})

	return r
}
