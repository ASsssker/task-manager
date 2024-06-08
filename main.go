package main

import "task-manager/cmd/server"

func main() {
	app := server.GetApp()
	app.RunServer()
}