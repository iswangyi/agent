package metrics

import (
	"agent/models"
	"agent/settings"
)


func SysMetrics() (L []*models.MetricValue) {
	L = append(L, models.GaugeValue("sysinfo.innerip", settings.IP()))
	return
}
