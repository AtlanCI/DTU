package cronjob

import (
	"DTU/pkg"
	"DTU/task/tuya"
)

func NewTuyaCron(cron *pkg.Crontab) {
	cron.AddByFunc("1", "* * * * *", func() {
		tuya.GetDeviceByUserId(pkg.GlobleId)
		tuya.GetDeviceInfo()
	})
}
