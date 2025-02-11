package main

import (
	"log"
	"net/http"
)

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("path %v method %v", r.URL.Path, r.Method)
	w.Write([]byte("hello"))
}

func main() {
	var s server
	http.ListenAndServe(":9999", &s)
}
