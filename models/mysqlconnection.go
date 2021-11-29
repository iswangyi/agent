package models

import (
	"agent/logger"
	"agent/settings"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Db() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/ops_gin?charset=utf8mb4&parseTime=True&loc=Local",
		settings.Config().Mysql.MysqlUser,
		settings.Config().Mysql.MysqlPasswd,
		settings.Config().Mysql.MysqlHost,
		settings.Config().Mysql.MysqlPort,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		_ = logger.StartupFatal("初始化数据库错误", err)
	}
	logger.StartupInfo("初始化数据库完成")
	return db
}
