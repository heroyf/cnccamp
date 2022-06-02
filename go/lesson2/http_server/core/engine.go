package core

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"
)

type handlerFunc func(c *Context)

type Engine struct {
	Router *Router
}

// New create engine object
func New() *Engine {
	return &Engine{Router: newRouter()}
}

// addRoute
func (e *Engine) addRoute(method string, url string, handler handlerFunc) {
	e.Router.addRoute(method, url, handler)
}

// Get add get route to route map
func (e *Engine) Get(url string, handler handlerFunc) {
	e.addRoute("GET", url, handler)
}

// Post add post route to route map
func (e *Engine) Post(url string, handler handlerFunc) {
	e.addRoute("POST", url, handler)
}

// Run start a http server
func (e *Engine) Run(addr string) (err error) {
	glog.Info(fmt.Sprintf("Starting http server: %s...", addr))
	return http.ListenAndServe(addr, e)
}

// ServerHttp handler func in map or return 404
func (e *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	context := newContext(response, request)
	e.Router.handle(context)
}
