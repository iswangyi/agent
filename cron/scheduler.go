package cron

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/robfig/cron/v3"
	"agent/logger"
	"time"
)


type SchedulerTask struct {
 	taskId string
 	cmd string
 	exec string
}

//Run xron会自动解析run方法并执行
func (task *SchedulerTask) Run()  {
	dt := &timestamp.Timestamp{Seconds: time.Now().Unix()}
	out ,exitCode,err := utils.
}

// CronStart 开启定时任务
func CronStart(taskId,cmd,exec,express string) cron.EntryID {
	sc := &SchedulerTask{
		taskId: taskId,
		cmd:    cmd,
		exec:   exec,
	}
	id,err := Xcron.AddJob(express,sc)
	if err !=nil {
		logger.JobInfo("addjob error:",err)
	}
	return id
}