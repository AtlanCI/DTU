package service

import (
	"DTU/configs/conf_v1"
	"DTU/internal/devices_health/biz"
	"DTU/pkg"
)

func StartHealthSend(config *conf_v1.Config) {
	Device := biz.NewDeviceHealthSend(config)
	pkg.DeviceHealthchannel.AddListener(Device)
}
