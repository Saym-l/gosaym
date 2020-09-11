package saym

import (
	"log"
	"net/http"
)

type router struct {
	Handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{Handlers: make(map[string]HandlerFunc)}
}
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %s -%s", method, pattern)
	key := method + "_" + pattern
	r.Handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.Handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOTFOUND:%S\n")
	}
}
