package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(response http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

// New Engine 工厂模式创建实例,返回对应实例指针
func New() *Engine {
	return &Engine{router: map[string]HandleFunc{}}
}

func (engine *Engine) addRouter(method string, path string, handler HandleFunc) {
	key := method + "-" + path
	fmt.Printf("add Route %v\n", key)
	engine.router[key] = handler
}

func (engine *Engine) GET(path string, handleFunc HandleFunc) {
	engine.addRouter("GET", path, handleFunc)
}

func (engine *Engine) POST(path string, handleFunc HandleFunc) {
	engine.addRouter("POST", path, handleFunc)
}

func (engine *Engine) Run(port string) (err error) {
	//可以强转
	//handler := (http.Handler)(engine)
	//return http.ListenAndServe(port, handler)

	return http.ListenAndServe(port, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %q %q \n", req.URL.Path, req.Method)
	}

}
