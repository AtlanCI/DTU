package service

import (
	"DTU/configs/conf_v1"
	"DTU/internal/devices_list/biz"
	"github.com/julienschmidt/httprouter"
)

func RegisterDeviceList(route *httprouter.Router, config *conf_v1.Config) {
	route.POST(config.ConnConfig.EcConfig.DeviceListInterface, biz.DeviceListHandler)
}
