//go:build linux || darwin || windows
// +build linux darwin windows

package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/heroyf/cnccamp/go/lesson12/service0/core"
	"github.com/heroyf/cnccamp/go/lesson12/service0/handler"
	"github.com/heroyf/cnccamp/go/lesson12/service0/metrics"
	"net/http"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("init http server...")
	// prometheus metrics register
	metrics.Register()

	engine := core.New()
	engine.Get("/", func(c *core.Context) {
		h := handler.GetHandler("RootHandler")
		h.GetHandleFunc()(c)
	})
	engine.Get("/delay", func(c *core.Context) {
		h := handler.GetHandler("LatencyHandler")
		h.GetHandleFunc()(c)
	})
	// answer 1 and 2 route
	engine.Get("/header", func(c *core.Context) {
		c.Json(http.StatusOK, c.W.Header())
	})
	engine.Get("/healthz", func(c *core.Context) {
		c.Html(http.StatusOK, "access to url[/healthz]")
	})
	engine.Get("/metrics", func(c *core.Context) {
		h := handler.GetHandler("MetricsHandler")
		h.GetHandleFunc()(c)
	})
	engine.Run(":8080")
}
