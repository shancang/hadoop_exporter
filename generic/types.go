package generic

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type CollectGenericMetricsForPrometheus struct {
	sync.Mutex

	Role     string
	Hostname string

	Namespace string

	Uri string
	HC  *http.Client

	CollectInterval    time.Duration
	CollectMetricsSets []prometheus.Metric

	ParseMetrics ParseMetrics
}

type ParseMetrics interface {
	ParseExporterStatus(ch chan<- prometheus.Metric, err error)
	ParseUniqueMetrics(chan prometheus.Metric, interface{})
}
