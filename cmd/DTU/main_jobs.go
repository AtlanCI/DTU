package main

import (
	"DTU/configs/conf_v1"
	"DTU/pkg"
	"DTU/task/cronjob"
)

var CronManager = pkg.NewCrontab()

func StartJobs(config *conf_v1.Config) {
	cronjob.NewAlleyesCron(CronManager, config)
	cronjob.NewConverCron(CronManager)
	cronjob.NewTuyaCron(CronManager)
}
