package collector

import (
	"agent/logger"
	"agent/metrics"
	"agent/models"
	"agent/settings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func Collect() {
	if !settings.Config().Transfer.Enable {
		return
	}
	if len(settings.Config().Transfer.Addr) == 0 {
		return
	}

	for _, v := range metrics.Mappers {
		if len(v.Fs) == 0 {
			return
		}
		collect(int64(v.Interval), v.Fs)

	}
}

func collect(sec int64, fns []func() []*models.MetricValue) {
	logger.ToMOCDebug("采集间隔sec:", sec)
	t := time.NewTicker(time.Second * time.Duration(sec))
	defer t.Stop()

	for {
		<-t.C
		hostname, err := settings.Hostname()
		if err != nil {
			continue
		}

		mvs := []*models.MetricValue{}
		//ignoreMetrics := settings.Config().IgnoreMetrics

		for _, fn := range fns {
			items := fn()
			if items == nil {
				continue
			}

			if len(items) == 0 {
				continue
			}

			for _, mv := range items {
				mvs = append(mvs, mv)
				// if b, ok := ignoreMetrics[mv.Metric]; ok && b {
				// 	continue
				// } else {
				// 	mvs = append(mvs, mv)
				// }
			}
		}

		now := time.Now().Unix()
		for j := 0; j < len(mvs); j++ {
			mvs[j].Step = sec
			mvs[j].Endpoint = hostname
			mvs[j].Timestamp = now
		}
		fmt.Println(mvs)
		//将指定内容写入到文件中
		out, err := json.Marshal(mvs)
		ss := string(out)
		err1 := ioutil.WriteFile("./output.txt", []byte(ss), 0666)
		log.Println(ss)
		if err1 != nil {
			fmt.Println("ioutil WriteFile error: ", err)
		}
		//		SendToTransfer(mvs)
	}
}
