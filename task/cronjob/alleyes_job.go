package cronjob

import (
	"DTU/pkg"
	"DTU/task/alleyes"
)

func NewAlleyesCron(cron *pkg.Crontab) {
	cron.AddByFunc("2", "* * * * *", func() {
		alleyes.GetIpcStatus("http://192.168.1.184:9008/v1/alleyes/syncbot/BoxStatus/get")
	})
}
