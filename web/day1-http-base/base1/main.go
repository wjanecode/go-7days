package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("serve start at :9999")
	err := http.ListenAndServe(":9999", nil) //阻塞调用
	if err != nil {
		fmt.Printf("error %q\n", err.Error())
	}
	fmt.Println("This will never be printed")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("received request for %s %s\n", req.Method, req.URL.Path)
	fmt.Fprintf(w, "Url.path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("received request for %s %s \n", req.Method, req.URL.Path)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
