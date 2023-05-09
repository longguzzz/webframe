package frame

import (
	"net/http"
)

type ServeFunc func(context *Context)
type Engine struct {
	*Router // 用指针
}

func NewEngine() *Engine {
	return &Engine{Router: newRouter(newTrie())}
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	e.Router.serveWith(context)
}
