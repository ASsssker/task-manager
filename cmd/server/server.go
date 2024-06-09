package server

import "net/http"

func (app *Applicaton) RunServer() error {
	
	err := http.ListenAndServe(app.ServerAddr, app.getRoutes())
	return err
}