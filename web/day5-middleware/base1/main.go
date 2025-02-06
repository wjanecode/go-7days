package main

import (
	"gee"
	"log"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", helloHandler)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hi", hiHandler)
	}

	r.Run(":9999")

}

func helloHandler(ctx *gee.Context) {
	log.Println("hello")
}
func hiHandler(ctx *gee.Context) {
	log.Println("say hi")
}
