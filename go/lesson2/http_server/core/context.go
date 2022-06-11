//go:build linux || darwin || windows
// +build linux darwin windows

package core

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	Url        string
	Method     string
	StatusCode int
}

// newContext create Context object
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		W:      w,
		R:      req,
		Url:    req.URL.Path,
		Method: req.Method,
	}
}

// SetStatus set resp status code
func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

// SetHeader set resp header
func (c *Context) SetHeader(key string, value string) {
	c.W.Header().Set(key, value)
}

// Json set json resp
func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}

// Data set resp data
func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.W.Write(data)
}

// Html return resp html
func (c *Context) Html(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	c.W.Write([]byte(html))
}
