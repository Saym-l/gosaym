package saym

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Write      http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Write:  w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Write.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Write.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Write.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Write)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Write, err.Error(), 500)
	}
}

func (c *Context) Date(code int, data []byte) {
	c.Status(code)
	c.Write.Write(data)
}
func (c *Context) Html(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Write.Write([]byte(html))
}
