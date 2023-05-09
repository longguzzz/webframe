package main

import (
	"net/http"
	"web/server/frame"
)

//	curl http://localhost:54321/

func main() {
	engine := frame.NewEngine()
	engine.Router.RigisterGET("/", func(c *frame.Context) {
		c.RespHtml(http.StatusOK, "<h1>root</h1>\n")
	})
	engine.Router.RigisterGET("/hello", func(c *frame.Context) {
		// /hello?name=test
		c.RespString(http.StatusOK, "para: %s", c.Query("name"))
	})
	engine.Router.RigisterPOST("/LOGIN", func(c *frame.Context) {
		c.RespJson(http.StatusOK, map[string]any{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.Run("localhost:54321")
}
