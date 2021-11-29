package metrics

import (
	models "agent/models"
	"agent/settings"
)

type FuncsAndInterval struct {
	Fs       []func() []*models.MetricValue
	Interval int
}

var Mappers []*FuncsAndInterval

func BuildMappers() {
	Interval := settings.Config().Transfer.Interval

	Mappers = []*FuncsAndInterval{
		//		&FuncsAndInterval{
		//			Fs: []func() []*models.MetricValue{
		//				SocketStatSummaryMetrics,
		//			},
		//			Interval: Interval,
		//		},
		&FuncsAndInterval{
			Fs: []func() []*models.MetricValue{
				LoadAvgMetrics,
			},
			Interval: Interval,
		},
	}
}
