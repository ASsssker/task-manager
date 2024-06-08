package server

import "net/http"

// Ping server connect
func (app *Applicaton) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}