package DTU

import (
	"DTU/configs/conf_v1"
	Event "DTU/internal/devices_event/service"
	Health "DTU/internal/devices_health/service"
	List "DTU/internal/devices_list/service"
	Recv "DTU/jobs/http_alleyes/service"
	Pulsar "DTU/jobs/pulsar_tuya/service"
	"DTU/pkg"
	"DTU/task/cronjob"
	"github.com/google/wire"
)

var AppProviderSet = wire.Build(conf_v1.NewConfig,
	pkg.NewHttpRouter,
	Event.StartEventSend,
	Health.StartHealthSend,
	List.RegisterDeviceList,
	Recv.RegisterAlleyesEvent,
	Pulsar.StartPulsar,
	pkg.NewCrontab,
	cronjob.NewAlleyesCron,
	cronjob.NewConverCron,
	cronjob.NewTuyaCron)
