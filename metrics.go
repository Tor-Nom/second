package second

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// metrics result
const (
	SUCCESS = "success"
	FAILED  = "failed"
)

// metrics level
const (
	LEVEL_UNKNOWN = "level_unknown"
	LEVEL_DEBUG   = "level_debug"
	LEVEL_INFO    = "level_info"
	LEVEL_WARN    = "level_warn"
	LEVEL_ERROR   = "level_error"
	LEVEL_FATAL   = "level_fatal"
)

var (
	NameSpace                      = "open_api"
	SubSystem                      = "open_api"
	errorMetricName                = "errorMetricName"
	counterErrorMetric             *prometheus.CounterVec
	policyGetUserInsureListService = InitPrometheusCounterVec(NameSpace, SubSystem, "policy_get_user_insure_list")
)

func init() {
	counterErrorMetric = InitPrometheusCounterVec(NameSpace, SubSystem, errorMetricName)
}

// 用于统计metrics
func Metrics(m *prometheus.CounterVec, result, stage, level, msgType string) {
	if m != nil {
		m.WithLabelValues(result, stage, level, msgType).Inc()
		return
	}
	// metrics 调用异常 上报
	counterErrorMetric.WithLabelValues(FAILED, "metrics.error", LEVEL_ERROR, "metrics.error").Inc()
}

// 用于初始化metrics
func InitPrometheusCounterVec(nameSpace, subSystem, name string) (baseCounter *prometheus.CounterVec) {
	baseCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: nameSpace,
		Subsystem: subSystem,
		Name:      name,
	}, []string{"result", "stage", `level`, "msgType"})
	return
}
