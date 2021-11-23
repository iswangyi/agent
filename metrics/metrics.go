package metrics

import model "agent/models"

type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

var Mappers []*FuncsAndInterval
