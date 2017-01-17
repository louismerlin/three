package main

import (
	"log"
	"net/http"

	"github.com/louismerlin/three/api/backend"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// websocket server
	server := backend.NewServer("/entry")
	go server.Listen()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
