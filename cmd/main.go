package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main()  {

	ch := make(chan os.Signal,1)
	signal.Notify(ch , os.Interrupt , os.Kill, syscall.SIGTERM)
	go func() {
		signalType := <-ch
		signal.Stop(ch)
		log.Info("退出")
		log.Info("收到os信号类型:", signalType)
	}()
	select {
	}
}

