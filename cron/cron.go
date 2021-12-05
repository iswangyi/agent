package cron

import (
	"agent/logger"
	"agent/settings"
	"errors"
	cron "github.com/robfig/cron/v3"
	"go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

var (
	Xcron *cron.Cron = cron.New(cron.WithParser(cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))
	client   *clientv3.Client
	initEtcd sync.Once
)

func init() {
	Xcron.Start()
}

func initScheduler() error {
	if settings.Config().Registrar.Enable {
		err := errors.New("已被初始化")
		logger.StartupDebug("cfg配置文件关闭etcd")
		return err
	} else {
		addr := settings.Config().Registrar.Addrs
		err := errors.New("etcd已被初始化")
		initEtcd.Do(func() {
			err = nil
			config := clientv3.Config{
				Endpoints:   addr,
				DialTimeout: 5 * time.Second,
			}
			// 建立客户端
			if client, err = clientv3.New(config); err != nil {
				logger.StartupFatal("链接etcd失败", err)
			} else {
				logger.StartupInfo("链接etcd成功")
			}
		})

		return err
	}
}
