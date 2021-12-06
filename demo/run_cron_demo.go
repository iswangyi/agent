package main

import (
	"fmt"
	QCron "github.com/robfig/cron/v3"
	"os/exec"
	"time"
)

var Qxron *QCron.Cron = QCron.New(QCron.WithParser(
	QCron.NewParser(
		QCron.Minute | QCron.Hour | QCron.Dom | QCron.Month | QCron.Dow)),
)

func init() {
	Qxron.Start()
}

type TestJob struct {
	taskId string
	cmd    string
}

func (this TestJob) Run() {
	dt := time.Now().Format("2006-01-02 15:00:04") //iso8601
	fmt.Println("yunxingzhong", dt)
	fmt.Println("JobId", this.taskId)
	c1, err := exec.Command("bash", "-c", this.cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(c1))

}

func cron() *QCron.EntryID {
	spec := "* * * * *"
	id, err := Qxron.AddJob(spec, &TestJob{
		taskId: "1",
		cmd:    "ls",
	})
	if err != nil {
		fmt.Println(err)
	}
	return &id
}

func main() {
	cron()
	select {}
}
