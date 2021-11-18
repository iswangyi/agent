package collector

import "agent/settings"

func Collect( )  {
	if !settings.Config().Transfer.Enable {
		return
	}
	if len( settings.Config().Transfer.Addr ) == 0 {
		return
	}

	for _,v := range



	
}
