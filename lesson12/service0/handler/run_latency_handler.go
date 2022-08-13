package handler

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/heroyf/cnccamp/go/lesson12/service0/core"
	"github.com/heroyf/cnccamp/go/lesson12/service0/metrics"
	"github.com/heroyf/cnccamp/go/lesson12/service0/utils"
	"net/http"
	"time"
)

func init() {
	Register("LatencyHandler", &LatencyHandler{})
}

type LatencyHandler struct{}

func (l LatencyHandler) GetHandleFunc() core.HandlerFunc {
	return func(c *core.Context) {
		glog.V(4).Info("enter latency handler....")
		timer := metrics.NewTimer()
		defer timer.ObserveTotal()
		// 添加0-2000ms的随机延时
		delay := utils.RandInt(0, 2000)
		time.Sleep(time.Millisecond * time.Duration(delay))
		timer.Delay = float64(delay) / 1000
		timer.ObserveDelay()
		glog.V(4).Infof("delay: %.3f s", timer.Delay)
		c.Html(http.StatusOK, fmt.Sprintf("delay: %.3f s", timer.Delay))
	}
}
