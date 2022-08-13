package handler

import (
	"github.com/golang/glog"
	"github.com/heroyf/cnccamp/go/lesson10/metric_http_server/core"
	"github.com/heroyf/cnccamp/go/lesson10/metric_http_server/utils"
	"io"
	"net/http"
	"strings"
	"time"
)

func init() {
	Register("RootHandler", &RootHandler{})
}

type RootHandler struct{}

func (l RootHandler) GetHandleFunc() core.HandlerFunc {
	return func(c *core.Context) {
		// 添加0-2000ms的随机延时
		delay := utils.RandInt(10, 20)
		time.Sleep(time.Millisecond * time.Duration(delay))
		io.WriteString(c.W, "========== Details of http request header: ==========\n")
		req, err := http.NewRequest("GET", "http://service2", nil)
		if err != nil {
			c.Html(http.StatusInternalServerError, "can not access to service1")
		}
		lowerCaseHeader := make(http.Header)
		for key, value := range c.R.Header {
			lowerCaseHeader[strings.ToLower(key)] = value
		}
		glog.Info("headers:", lowerCaseHeader)
		req.Header = lowerCaseHeader
		c.Html(http.StatusOK, "access to root")
	}
}
