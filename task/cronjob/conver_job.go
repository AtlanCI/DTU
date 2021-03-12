package cronjob

import (
	"DTU/pkg"
	"DTU/task/conversion"
)

func NewConverCron(cron *pkg.Crontab) {
	cron.AddByFunc("3", "* * * * *", func() {
		device := conversion.NewDeviceCache()
		pkg.DeviceCacheChannel.AddListener(&device)
		device.ConversionHealth()
		device.ConversionEvent()
	})
}
