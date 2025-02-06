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

	v1 := r.Group("/v1")

	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
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
