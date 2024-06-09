package server

import "net/http"

func (app *Applicaton) RunServer() error {

	app.InfoLog.Printf("starting server on %s", app.ServerAddr)
	err := http.ListenAndServe(app.ServerAddr, app.getRoutes())
	return err
}
