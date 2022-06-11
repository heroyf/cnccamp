package main

import "C"
import (
	"flag"
	"github.com/golang/glog"
	core "github.com/heroyf/cnccamp/go/lesson2/http_server/core"
	"net/http"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("init http server...")
	engine := core.New()
	engine.Get("/", func(c *core.Context) {
		c.Html(http.StatusOK, "access to root")
	})
	// answer 1 and 2 route
	engine.Get("/header", func(c *core.Context) {
		c.Json(http.StatusOK, c.W.Header())
	})
	engine.Get("/healthz", func(c *core.Context) {
		c.Html(http.StatusOK, "access to url[/healthz]")
	})
	engine.Run(":8080")
}
