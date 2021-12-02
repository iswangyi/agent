package main

import (
	"agent/collector"
	"agent/http"
	"agent/logger"
	"agent/metrics"
	"agent/models"
	"agent/settings"
	"fmt"
	goHttp "net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	logger.Init()

	//处理程序接收到的系统signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		signalType := <-ch
		signal.Stop(ch)
		logger.Info("退出")
		logger.Info("收到os信号类型:", signalType)
	}()
	if len(os.Args) != 2 {
		fmt.Println("Start error, [start|stop|version]")
		os.Exit(1)
	}
	//初始化配置文件
	settings.LoadConfiguration()
	//加载数据库
	models.Db()
	//初始化监控项
	metrics.BuildMappers()
	//执行监控
	go collector.Collect()
	if strings.ToLower(os.Args[1]) == "main" {
		http.Start()
		go func() {
			//利用net/http具备守护进程的能力
			goHttp.ListenAndServe("0.0.0.0:10028", nil)
		}()
	}
	settings.HandleControl(os.Args[1])
}
