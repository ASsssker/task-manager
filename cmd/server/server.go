package server

import "net/http"

func (app *Applicaton) RunServer() error {
	
	err := http.ListenAndServe(app.ServerAddr, app.geRoutes())
	return err
}