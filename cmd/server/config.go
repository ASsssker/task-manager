package server

import "flag"

type config struct {
	ServerAddr string
}


func getConfig() (*config, error) {
	var conf config

	flag.StringVar(&conf.ServerAddr, "a", "localhost:8080", "Server addres in the format ip:port")

	flag.Parse()

	return &conf, nil
}