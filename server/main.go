package main

import (
	"net/http"
	"web/server/frame"
)

//	curl http://localhost:54321/

func main() {
	engine := frame.NewEngine()
	// engine.Router.RigisterGET("/", func(c *frame.RequestContext) {
	// 	c.RespHtml(http.StatusOK, "<h1>root</h1>\n")
	// })
	// engine.Router.RigisterGET("/hello", func(c *frame.RequestContext) {
	// 	// /hello?name=test
	// 	c.RespString(http.StatusOK, "para: %s", c.Query("name"))
	// })
	// engine.Router.RigisterPOST("/LOGIN", func(c *frame.RequestContext) {
	// 	c.RespJson(http.StatusOK, map[string]any{
	// 		"username": c.PostForm("username"),
	// 		"password": c.PostForm("password"),
	// 	})
	// })

	g := engine.NewSubGroup("login")
	// curl "http://localhost:54321/login" -X POST -d 'username=test&password=1234'
	// {"password":"1234","username":"test"}
	g.RegisterPOST(func(c *frame.RequestContext) {
		c.RespJson(http.StatusOK, map[string]any{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	// curl "http://localhost:54321/login/second"
	g.NewSubGroup("second").RegisterGET(func(c *frame.RequestContext) {
		c.RespString(http.StatusOK, "para: %s", c.Query("name"))
	})

	engine.Run("localhost:54321")
}
