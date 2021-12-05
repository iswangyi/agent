package cron

import "agent/settings"

func GetTaskListUrl() string {
	return "/cron/jobs/" + settings.IP()
}
func GetTaskCronIDUrl() string {
	return "/cron/cronid/jobs" + settings.IP() + "/"
}
func GetTaskIDUrl() string {
	return GetTaskListUrl() + "/"
}
