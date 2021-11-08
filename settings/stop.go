package settings

import (
	"agent/logger"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func processInfo(pid int) (string,error)  {
	p,err := process.NewProcess(int32(pid))
	if err != nil {
		msg := fmt.Sprintf("Can't read process info")
		logger.StartupInfo(msg,"pid=",pid)
		return msg,err
	}
	if v,err := p.Cmdline();err==nil{
		msg := fmt.Sprintf("cmd +args:")
		logger.FatalInfo(msg)
		return msg,err
	}
	return "",nil
}

func StopHandle  {
	if _,err := os.Stat(config.Pid);err == nil{
		data,err := ioutil.ReadFile(Config().Pid)
		if err != nil {
			fmt.Println("agent not running")
			logger.FatalInfo("Not Running")
		}
		logger.StartupInfo(string(data))
		ProcessId,err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("Unable to read and praise id")
			logger.FatalInfo("get pid error")
		}
		//二次校验
		process,err := os.FindProcess(ProcessId)
		if err != nil {
			logger.FatalInfo("get pid error")
		}
		os.Remove(Config().Pid)
		PidInfo ,_ := processInfo(ProcessId)
		if strings.Contains(PidInfo) {

		}

	}

}