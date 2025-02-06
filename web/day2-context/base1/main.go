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
	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
}

func helloHandler(c *gee.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func loginHandler(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
