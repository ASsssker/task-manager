package main

import "task-manager/cmd/server"

func main() {
	err := server.RunServer()
	if err != nil {
		panic("Server error")
	}
}