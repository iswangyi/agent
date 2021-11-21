package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	logger  *log.Logger
	intiLog sync.Once
)

func Init() error {
	fmt.Println("init log")
	//设置日志格式为json
	err := errors.New("初始化")
	logger = log.New()
	intiLog.Do(func() {
		err = nil
		logger.Formatter = &log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2021-12-01 12:22:22",
		}
		var filename = "./logfile.log"
		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		logger.Out = f
		logger.Level = log.DebugLevel
	})

	return err
}

func WithField(key string, value interface{}) *log.Entry {
	return log.WithField(key, value)
}

func Debug(args ...interface{}) {
	log.Info(args)
}
func Info(args ...interface{}) {
	log.Info(args)
}
func StartupInfo(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "startup").Info(msg...)
		return err
	}
	WithField("key", "startup").Info(msg...)
	return nil
}
func Fatal(args ...interface{}) {
	log.Fatal(args)
}
func FatalInfo(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "Fatal error").Fatal(msg...)
		return err
	}
	WithField("key", "Fatal error").Fatal(msg...)
	return nil
}
