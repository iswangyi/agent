package models

import (
	"agent/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type SystemCollector struct {
	CpuMetric        *prometheus.Desc
	ServerInfoMetric *prometheus.Desc
}

func NewSystemCollector() *SystemCollector {
	variableLabels := []string{"zheng"}
	constLabels := make(map[string]string)
	constLabels["env"] = "prod"
	return &SystemCollector{
		CpuMetric:        prometheus.NewDesc("CpuLoad", "cpu load info", variableLabels, constLabels),
		ServerInfoMetric: prometheus.NewDesc("network", "network speed", variableLabels, constLabels),
	}
}

func (collect *SystemCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collect.CpuMetric
	ch <- collect.ServerInfoMetric

}

func (collect *SystemCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64

	if 1 == 1 {
		metricValue = 1
	}
	ch <- prometheus.MustNewConstMetric(collect.CpuMetric, prometheus.GaugeValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collect.ServerInfoMetric, prometheus.CounterValue, metricValue, "kk")
}

func PromCollect() {
	logger.StartupInfo("start collect ")
	sys := NewSystemCollector()
	prometheus.MustRegister(sys)
	promhttp.Handler()
}
