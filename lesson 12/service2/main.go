//go:build linux || darwin || windows
// +build linux darwin windows

package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/heroyf/cnccamp/go/lesson10/metric_http_server/core"
	"github.com/heroyf/cnccamp/go/lesson10/metric_http_server/handler"
	"github.com/heroyf/cnccamp/go/lesson10/metric_http_server/metrics"
	"net/http"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("init http server...")
	// prometheus metrics register
	metrics.Register()

	engine := core.New()
	engine.Get("/", func(c *core.Context) {
		c.Html(http.StatusOK, "access to root")
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
