package main

import (
	"log"
	"net/http"
)

func main() {
	s := server.New()

	log.Fatal(http.ListenAndServe(":1880", s.Router()))
}
