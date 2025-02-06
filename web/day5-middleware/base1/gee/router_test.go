package gee

import (
	"log"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test pattern fail")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	//r.addRoute("GET", "/hello/1", nil)

	n, p := r.getRoute("GET", "/hello/key")
	node := r.getRoutes("GET")
	log.Printf("%v", node)
	log.Printf("%v, %v", n, p)
	if n == nil {
		t.Fatal("pattern return nil")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

}
