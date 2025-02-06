package gee

import (
	"log"
	"testing"
)

func TestNextGroup(t *testing.T) {
	r := New()
	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2 but now" + v2.prefix)
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2 but now " + v3.prefix)
	}
}

func TestUseMiddleware(t *testing.T) {
	r := New()
	v1 := r.Group("/v1")

	v1.Use(middle1)
	r.Run(":9999")
}

func middle1(c *Context) {
	log.Printf("%v %v \n", c.Method, c.Path)
}
