package main

import (
	"crowns/app/server"
	"crowns/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Fatal(
		RunServer(),
	)
}

func RunServer() error {
	s := &server.SimpleServer{}
	s.Init(config.Load())

	http.Handle("/", s.Router())

	fmt.Println("Server starting on http://0.0.0.0:8081")
	return http.ListenAndServe("0.0.0.0:8081", nil)
}
