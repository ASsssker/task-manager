package main

import (
	"task-manager/cmd/server"
)

func main() {
	app, err := server.GetApp()
	if err != nil {
		panic(err)
	}
	app.RunServer()
}