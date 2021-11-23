package metrics

import "agent/models"

//SocketStatSummaryMetrics TCP链接
func SocketStatSummaryMetrics() []*models.MetricValue {
	return []*models.MetricValue{
		&models.MetricValue{
			Endpoint:  "",
			Metric:    "",
			Value:     nil,
			Step:      0,
			Type:      "",
			Tags:      "",
			Timestamp: 0,
		},
	}
}
