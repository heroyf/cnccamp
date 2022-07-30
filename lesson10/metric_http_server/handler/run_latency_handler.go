package handler

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/heroyf/cnccamp/go/lesson10/metric_http_server/core"
	"github.com/heroyf/cnccamp/go/lesson10/metric_http_server/metrics"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	Register("LatencyHandler", &LatencyHandler{})
}

type LatencyHandler struct{}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func (l LatencyHandler) GetHandleFunc() core.HandlerFunc {
	return func(c *core.Context) {
		glog.V(4).Info("enter latency handler....")
		timer := metrics.NewTimer()
		defer timer.ObserveTotal()
		// 添加0-2000ms的随机延时
		delay := randInt(0, 2000)
		time.Sleep(time.Millisecond * time.Duration(delay))
		glog.V(4).Infof("delay: %d ms", delay)
		c.Html(http.StatusOK, fmt.Sprintf("delay: %d ms", delay))
	}
}
