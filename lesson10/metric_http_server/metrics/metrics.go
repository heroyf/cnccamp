package metrics

import (
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	Namespace = "http_server"
)

var (
	latency = createExecutionTimeMetrics(Namespace, "http server time spent.")
)

// ExecutionTimer 执行计时器
type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
	end   time.Time
}

// createExecutionTimeMetrics 创建http server执行时间指标
func createExecutionTimeMetrics(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "execution_latency_seconds",
			Help:      help,
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"steps"},
	)
}

// Register 注册指标
func Register() {
	err := prometheus.Register(latency)
	if err != nil {
		glog.Errorf("prometheus register error: %v", err)
	}
}

// NewTimer 启动一个histogramVec的计时器
func NewTimer() *ExecutionTimer {
	return NewExecutionTime(latency)
}

// NewExecutionTime 启动一个histogramVec的计时器
func NewExecutionTime(histo *prometheus.HistogramVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: histo,
		start: now,
		end:   now,
	}
}

// ObserveTotal 计算使用总时长
func (t *ExecutionTimer) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}
