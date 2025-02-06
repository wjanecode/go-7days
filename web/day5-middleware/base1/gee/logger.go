package gee

import (
	"log"
	"time"
)

func Logger() HandleFunc {
	return func(ctx *Context) {
		t := time.Now()
		ctx.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
