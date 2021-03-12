package service

import (
	"DTU/configs/conf_v1"
	"DTU/internal/devices_event/biz"
	"DTU/pkg"
)

func StartEventSend(config *conf_v1.Config) {
	Event := biz.NewDeviceEventSend(config)
	pkg.DeviceEventChannel.AddListener(Event)
}
