package main

import (
	"DTU/configs/conf_v1"
	Event "DTU/internal/devices_event/service"
	Health "DTU/internal/devices_health/service"
)

func StartUpDataToEc(config *conf_v1.Config) {
	Event.StartEventSend(config)
	Health.StartHealthSend(config)
}
