package main

import "C"
import (
	"flag"
	"github.com/golang/glog"
	"github.com/heroyf/cnccamp/lesson1/http_server/core"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("init http server...")
	engine := core.New()
	// answer 1 and 2 route
	engine.Get("/", func(c *core.Context) {
		c.Json(200, c.W.Header())
	})
	engine.Get("/healthz", func(c *core.Context) {
		c.Html(200, "access to url[/healthz]")
	})
	engine.Run(":8080")
}
