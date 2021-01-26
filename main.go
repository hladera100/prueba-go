package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":1880", nil))
}
