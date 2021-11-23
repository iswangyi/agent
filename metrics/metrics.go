package metrics

import (
	model "agent/models"
	"agent/settings"
)

type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

var Mappers []*FuncsAndInterval

func BuildMappers() {
	Interval := settings.Config().Transfer.Interval

	Mappers = []*FuncsAndInterval{
		&FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				CpuMetrics,
			},
			Interval: Interval,
		},
		&FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				SocketStatSummaryMetrics,
			},
			Interval: Interval,
		},
	}
}
