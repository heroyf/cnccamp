package handler

import (
	"github.com/heroyf/cnccamp/go/lesson12/service0/core"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	Register("MetricsHandler", &MetricsHandler{})
}

type MetricsHandler struct{}

func (l MetricsHandler) GetHandleFunc() core.HandlerFunc {
	h := promhttp.Handler()

	return func(c *core.Context) {
		h.ServeHTTP(c.W, c.R)
	}
}
