package settings

import (
	"agent/logger"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
	"os/exec"
	"strconv"
)

func FindCheckProcess() ([]*process.Process, error) {
	return process.Processes()
}

//进程无法感知自己，采用第三方库检测
func getAgentList() []string {
	list, _ := FindCheckProcess()
	var agentNum []string
	for _, v := range list {
		info, _ := v.Name()
		if info == "agent_osx" {
			agentNum = append(agentNum, info)
		}
	}
	return agentNum
}

// StartHandle 处理程序启动
func StartHandle() {
	//检查agent 是否已经运行
	if _, err := os.Stat(Config().Pid); err == nil {
		logger.Fatal("start up failed")
		os.Exit(1)
	}
	agentList := getAgentList()
	if len(agentList) > 1 {
		fmt.Println(agentList)
		logger.Fatal("start up failed, agent process exists")
		os.Exit(1)
	}
	cmd := exec.Command(os.Args[0], "main")
	err := cmd.Start()
	if err != nil {
		logger.StartupInfo("start cmd.start", err)
	}
	logger.Info("init config successful")
	logger.Info("version", GetVersion())
	logger.Info("pid is:", cmd.Process.Pid)
	savePID(cmd.Process.Pid)
	os.Exit(0)
}

//保存pid信息
func savePID(pid int) {
	Config().Pid = strconv.Itoa(pid)
	logger.StartupInfo("已保存pid至cfg.json")
	file, err := os.Create(Config().Pid)
	_, err = file.WriteString(strconv.Itoa(pid))
	if err != nil {
		logger.StartupInfo("save pid failed")
	}
	if err != nil {
		fmt.Println("pid file not exists")
		logger.FatalInfo("没有pid")
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	file.Sync()
}
