// +build wireinject

package biz

import (
	"DTU/internal/devices_list/data"
	"github.com/google/wire"
)

var DeviceDataSet = wire.NewSet(data.NewDeviceListCache, wire.Bind(new(CacheList), new(*data.DeviceListCache)), NewCacheBiz)

func initDeviceDataSet() *CacheBiz {
	wire.Build(DeviceDataSet)
	return &CacheBiz{}
}
