package gee

import "net/http"

type router struct {
	handlers map[string]HandlerFun
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFun)}
}

func (r *router) addRouter(method string, path string, handler HandlerFun) {
	key := method + "-" + path
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handlers, ok := r.handlers[key]; ok {
		handlers(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND")
	}

}
