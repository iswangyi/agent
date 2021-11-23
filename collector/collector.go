package collector

import (
	"agent/metrics"
	"agent/settings"
	"go/format"
	"time"
)

func Collect( )  {
	if !settings.Config().Transfer.Enable {
		return
	}
	if len( settings.Config().Transfer.Addr ) == 0 {
		return
	}

	for _,v := range metrics.Mappers {
		go
	}
}

func collect(sec int64,fns *metrics.FuncsAndInterval)  {
	t := time.NewTicker(time.Second * time.Duration(sec)).C
	for {
		<-t
		switch expr {
		
		}	
			
		
	}
	







}
