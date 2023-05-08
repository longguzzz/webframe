package frame

import (
	"net/http"
)

type ServeFunc func(context *Context)
type Engine struct {
	router *router // 用指针
}

func NewEngine() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) GET(pattern string, serveF ServeFunc) {
	e.router.addRoute("GET", pattern, serveF)
}

func (e *Engine) POST(pattern string, serveF ServeFunc) {
	e.router.addRoute("POST", pattern, serveF)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	e.router.serveAll(context)
}
