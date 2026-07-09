package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	log.Println("bounty-api listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
