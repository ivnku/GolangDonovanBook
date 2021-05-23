package main

import (
	"log"
	"net/http"
	lissajous "donovanBook/gif"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}