package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.Default()

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{
		Name: "woo",
		Age:  18,
	}
	stu2 := &student{
		Name: "Jack",
		Age:  20,
	}

	r.GET("/", func(context *gee.Context) {
		context.HTML(
			http.StatusOK, "index.tmpl", gee.H{
				"title":  "geek",
				"stuArr": [2]*student{stu1, stu2},
			})
	})

	r.GET("/panic", func(context *gee.Context) {
		arr := []string{"hello"}
		context.String(http.StatusOK, arr[100])
	})
	r.Run(":9999")
}
