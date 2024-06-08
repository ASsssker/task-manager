package server

import "log"



type Applicaton struct {
	*config
}

func GetApp() *Applicaton {
	var app Applicaton
	
	conf, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}
	app.config = conf

	return &app
}

func (app *Applicaton) Run() error{
	err := app.RunServer()
	return err
}