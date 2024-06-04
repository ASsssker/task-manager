package server

import "net/http"

func RunServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK!"))
	})

	err := http.ListenAndServe(":8080", nil)
	return err
}