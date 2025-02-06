package gee

import (
	"net/http"
)

type HandlerFun func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) addRouter(method string, path string, handler HandlerFun) {
	e.router.addRouter(method, path, handler)
}

func (e *Engine) GET(path string, handler HandlerFun) {
	e.addRouter("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFun) {
	e.addRouter("POST", path, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

func (e *Engine) Run(port string) error {
	return http.ListenAndServe(port, e)
}
