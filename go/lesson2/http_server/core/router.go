package core

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/heroyf/cnccamp/lesson2/http_server/version"
	"net"
	"net/http"
	"os"
	"strings"
)

type Router struct {
	Handlers map[string]handlerFunc
}

func newRouter() *Router {
	return &Router{Handlers: make(map[string]handlerFunc)}
}

// addRoute add url route
func (r *Router) addRoute(method string, url string, handler handlerFunc) {
	r.Handlers[method+"-"+url] = handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Url
	if handler, ok := r.Handlers[key]; ok {
		writeRespHeader(c.W, c.R)
		handler(c)
		ip, err := getClientIp(c.R)
		if err != nil {
			glog.Errorf("get client ip error: %v", err)
		}
		glog.V(2).Infof("url: [%s] - client: [%s], returnCode: [%s]", c.Url, ip, c.StatusCode)
	} else {
		c.Html(http.StatusNotFound, fmt.Sprintf("404 NOT FOUND: %s\n", c.Url))
	}
}

// writeRespHeader write resp header
func writeRespHeader(response http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		response.Header().Set(k, fmt.Sprint(v))
	}
	// answer 2: write sys env[VERSION]
	sysVersion := os.Getenv("VERSION")
	if sysVersion == "" {
		sysVersion = version.Version
	}
	response.Header().Set("Sys-Version", sysVersion)
}

// getClientIp get client ip
func getClientIp(r *http.Request) (string, error) {
	// get X-Real-IP  header
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	// get X-Forward-For header
	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	// finally get remote addr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}
