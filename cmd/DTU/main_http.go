package main

import (
	"DTU/configs/conf_v1"
	DeviceList "DTU/internal/devices_list/service"
	AlleyesRecv "DTU/jobs/http_alleyes/service"
	"DTU/pkg"
	"net/http"
)

var Router = pkg.NewHttpRouter()

func StartHttpServer(config *conf_v1.Config) {
	DeviceList.RegisterDeviceList(Router, config)
	AlleyesRecv.RegisterAlleyesEvent(Router, config)
	http.ListenAndServe(config.ListenPort, Router)
}
