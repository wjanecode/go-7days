package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
)

func main() {

	engine := gee.New()
	engine.GET("/", indexHandler)
	engine.GET("/hello", helloHandler)
	log.Fatal(engine.Run(":9999"))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "url = %q method=%q \n", req.URL.Path, req.Method)
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "header %q = %q \n", k, v)
	}
}
