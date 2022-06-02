package main

import "C"
import (
	"flag"
	"github.com/golang/glog"
	core "github.com/heroyf/cnccamp/lesson2/http_server/core"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("init http server...")
	engine := core.New()
	engine.Get("/", func(c *core.Context) {
		c.Html(200, "access to root")
	})
	// answer 1 and 2 route
	engine.Get("/header", func(c *core.Context) {
		c.Json(200, c.W.Header())
	})
	engine.Get("/healthz", func(c *core.Context) {
		c.Html(200, "access to url[/healthz]")
	})
	engine.Run(":8080")
}
