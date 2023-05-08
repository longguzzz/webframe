package frame

import (
	"log"
	"net/http"
)

type router struct {
	ServeFuncMap map[string]ServeFunc
}

func newRouter() *router {
	return &router{ServeFuncMap: make(map[string]ServeFunc)}
}

func (r *router) addRoute(method string, pattern string, serveF ServeFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	r.ServeFuncMap[key] = serveF
}

// func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	context := NewContext(w, r)

//		key := r.Method + "-" + r.URL.Path
//		if serveF, ok := e.router.ServeFuncMap[key]; ok {
//			serveF(context)
//		} else {
//			context.RespString(http.StatusNotFound, "404 NOT FOUND: %s\n", r.URL)
//		}
//	}
func (r *router) serveAll(c *Context) {
	key := c.Request.Method + "-" + c.Request.URL.Path
	if serveF, ok := r.ServeFuncMap[key]; ok {
		serveF(c)
	} else {
		c.RespString(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Request.URL)
	}
}
