package settings

import (
	"agent/logger"
	"os"
)

func getAgentList() []string {
	list,_ := FindCheckProcess()

}

func StartHandle()  {
	//检查agent 是否已经运行
	if _,err := os.Stat(Config().Pid); err ==nil {
		logger.Fatal("start up failed")
		os.Exit(1)
	}
}
