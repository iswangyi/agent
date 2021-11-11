package logger

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	logger  *log.Logger
	intiLog sync.Once
)

func Init() error {

	err := errors.New("logger已被初始化")
	intiLog.Do(func() {
		err = nil
		logger = log.New()
		logger.Formatter = &log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
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
	logger.Debug(args)
}
func Info(args ...interface{}) {
	logger.Info(args)
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
	logger.Fatal(args)
}
func FatalInfo(msg ...interface{}) error {
	if err := Init(); err != nil {
		WithField("key", "Fatal error").Fatal(msg...)
		return err
	}
	WithField("key", "Fatal error").Fatal(msg...)
	return nil
}
