package main

import (
	"agent/collector"
	"agent/http"
	log "agent/logger"
	"agent/metrics"
	"agent/settings"
	"fmt"
	goHttp "net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	log.Init()

	//处理程序接收到的系统signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		signalType := <-ch
		signal.Stop(ch)
		log.Info("退出")
		log.Info("收到os信号类型:", signalType)
	}()
	if len(os.Args) != 2 {
		fmt.Println("Start error, [start|stop|version]")
		os.Exit(1)
	}
	settings.LoadConfiguration()
	settings.InitLocalIp()
	metrics.BuildMappers()
	collector.Collect()
	if strings.ToLower(os.Args[1]) == "main" {
		http.Start()
		go func() {
			//利用net/http具备守护进程的能力
			goHttp.ListenAndServe("0.0.0.0:10028", nil)
		}()
	}
	settings.HandleControl(os.Args[1])
}
