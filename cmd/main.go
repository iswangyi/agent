package main

import (
	"agent/http"
	"agent/settings"
	"fmt"
	log "github.com/sirupsen/logrus"
	goHttp "net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {

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
	if strings.ToLower(os.Args[1]) == "main" {
		go func() {
			fmt.Println("启动main程序2:", os.Args[1])
			//利用net/http具备守护进程的能力
			goHttp.ListenAndServe("0.0.0.0:10028", nil)
		}()
		fmt.Println("启动main程序3:", os.Args[1])
		http.Start()
	}
	settings.HandleControl(os.Args[1])
	time.Sleep(10 * time.Second)
}
