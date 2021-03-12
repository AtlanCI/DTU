package service

import (
	"DTU/internal/devices_list/biz"
	"github.com/julienschmidt/httprouter"
)

func RegisterDeviceList(route *httprouter.Router) {
	route.POST("/ha/device/list", biz.DeviceListHandler)
}
