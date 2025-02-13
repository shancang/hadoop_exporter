package hadoop

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

// "Hadoop:service=NameNode,name=RetryCache.NameNodeRetryCache"
func (c *Collect) parseNameNodeRetryCache(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "CacheHit", "CacheCleared", "CacheUpdated":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "namenode_retry_cache", metricsName),
					strings.Join([]string{c.Namespace, "namenode retry cache", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		}
	}
}
