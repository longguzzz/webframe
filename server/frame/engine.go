package webframe

import (
	"fmt"
	"log"
	"net/http"
)

type ServeF func(w http.ResponseWriter, r *http.Request)
type Engine struct {
	router map[string]ServeF
}

func NewEngine() *Engine {
	return &Engine{router: make(map[string]ServeF)}
}

func (e *Engine) GET(pattern string, serveF ServeF) {
	e.addRoute("GET", pattern, serveF)
}

func (e *Engine) POST(pattern string, serveF ServeF) {
	e.addRoute("POST", pattern, serveF)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) addRoute(method string, pattern string, serveF ServeF) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	e.router[key] = serveF
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if serveF, ok := e.router[key]; ok {
		serveF(w, r)
	} else {
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
		if err != nil {
			ErrBacktrace(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
