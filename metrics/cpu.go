package metrics

import "agent/models"

func CpuMetrics() []*models.MetricValue {

	return models.MetricValue{
		Endpoint:  "",
		Metric:    "",
		Value:     nil,
		Step:      0,
		Type:      "",
		Tags:      "",
		Timestamp: 0,
	}
}
