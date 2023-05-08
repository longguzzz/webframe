package frame

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 使用Context进行更细粒度的控制，同时维护Writer的信息，比如Header
type Context struct {
	Writer  http.ResponseWriter // 接口
	Request *http.Request       // 指向结构体

	Path       string
	Method     string
	StatusCode int
}

// TODO: 返回指针的理由
func NewContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:     writer,
		Request:    req,
		Path:       req.URL.Path,
		Method:     req.Method,
		StatusCode: 0,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 数据封装
func (c *Context) RespString(code int, format string, values ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// TODO: string与[]byte类型转换的细节
func (c *Context) RespHtml(code int, html string) {
	c.SetHeader("Content-Type","text/html")
	c.SetStatus(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) RespJson(code int, obj any) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError) // 500
	}
}

func (c *Context) RespData(code int,data []byte) {
	c.SetStatus(code)
	c.Writer.Write(data)
}

