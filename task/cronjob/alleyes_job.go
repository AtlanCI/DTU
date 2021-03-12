package cronjob

import (
	"DTU/configs/conf_v1"
	"DTU/pkg"
	"DTU/task/alleyes"
)

func NewAlleyesCron(cron *pkg.Crontab, config *conf_v1.Config) {
	cron.AddByFunc("2", "* * * * *", func() {
		alleyes.GetIpcStatus(config.ConnConfig.AlleyesConfig.IPCStatusGettingInterface)
	})
}
