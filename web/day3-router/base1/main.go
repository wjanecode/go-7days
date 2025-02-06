package main

import (
	"gee"
	"net/http"
)

func main() {

	r := gee.New()
	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)
	r.POST("/login", loginHandler)
	r.Run(":9999")
}

func indexHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "welcome")
}

func helloHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "hello")
}

func loginHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "login success")
}
